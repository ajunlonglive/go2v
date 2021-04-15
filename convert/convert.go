package convert

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"strings"
)

func Convert(filename string) (string, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return "", err
	}
	cfg := &types.Config{Importer: importer.Default()}
	info := types.Info{}
	_, err = cfg.Check(filename, fset, []*ast.File{file}, &info)
	if err != nil {
		return "", err
	}

	c := &convert{file: file, info: &info}
	c.convert()
	return c.out.String(), nil
}

type convert struct {
	file   *ast.File
	info   *types.Info
	out    strings.Builder
	indent int
}

func (c *convert) write(s string) {
	c.out.WriteString(s)
}

func (c *convert) writeln(s string) {
	c.out.WriteString(s + "\n")
	if c.indent < 13 {
		tabs := [13]string{"", "\t", "\t\t", "\t\t\t", "\t\t\t\t", "\t\t\t\t\t", "\t\t\t\t\t\t", "\t\t\t\t\t\t\t", "\t\t\t\t\t\t\t\t", "\t\t\t\t\t\t\t\t\t", "\t\t\t\t\t\t\t\t\t\t", "\t\t\t\t\t\t\t\t\t\t\t", "\t\t\t\t\t\t\t\t\t\t\t\t"}
		c.out.WriteString(tabs[c.indent])
	} else {
		for i := 0; i < c.indent; i++ {
			c.out.WriteRune('\t')
		}
	}
}

func (c *convert) convert() {
	c.write("module ")
	c.expr(c.file.Name)

	c.decls(c.file.Decls)
}

func (c *convert) decls(decls []ast.Decl) {
	for _, decl := range decls {
		c.write("\n")
		c.decl(decl)
	}
}

func (c *convert) decl(_decl ast.Decl) {
	switch decl := _decl.(type) {
	case *ast.FuncDecl:
		if decl.Recv != nil {
			panic("cannot do recievers yet")
		}
		c.write("fn ")
		c.expr(decl.Name)
		c.write("() ")
		c.block(decl.Body)
	case *ast.GenDecl:
		if decl.Tok == token.IMPORT {
			for _, spec := range decl.Specs {
				c.spec(spec)
			}
		} else {
			panic(fmt.Sprintf("unknown generic decleration: %v", decl.Tok))
		}
	default:
		panic(fmt.Sprintf("unknown type: %T", _decl))
	}
}

func (c *convert) spec(_spec ast.Spec) {
	switch spec := _spec.(type) {
	case *ast.ImportSpec:
		if spec.Path.Value == "\"fmt\"" {

		} else {
			panic("unknown import: " + spec.Path.Value)
		}
	default:
		panic(fmt.Sprintf("unknown type: %T", _spec))
	}
}

func (c *convert) block(block *ast.BlockStmt) {
	c.indent++
	c.writeln("{")
	for i, stmt := range block.List {
		if i == len(block.List)-1 {
			c.indent--
		}
		c.stmt(stmt)
	}

	c.writeln("}")
}
