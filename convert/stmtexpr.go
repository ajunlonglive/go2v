package convert

import (
	"fmt"
	"go/ast"
	"go/token"
)

func (c *convert) stmt(_stmt ast.Stmt) {
	switch stmt := _stmt.(type) {
	case *ast.ExprStmt:
		c.expr(stmt.X)
	default:
		panic(fmt.Sprintf("unknown statement: %T", _stmt))
	}
	c.writeln("")
}

func (c *convert) expr(_expr ast.Expr) {
	switch expr := _expr.(type) {
	case *ast.CallExpr:
		c.map_fn_expr(expr.Fun)
		c.write("(")
		for i, arg := range expr.Args {
			c.expr(arg)
			if i != len(expr.Args)-1 {
				c.write(", ")
			}
		}
		c.write(")")
	case *ast.SelectorExpr:
		c.expr(expr.X)
		c.write(".")
		c.expr(expr.Sel)
	case *ast.Ident:
		fmt.Printf("obj: %v, name: %v\n", expr.Obj, expr.Name)
		c.write(expr.Name)
	case *ast.BasicLit:
		switch expr.Kind {
		case token.INT:
			c.write(expr.Value)
		case token.FLOAT:
			c.write(expr.Value)
		case token.IMAG:
			panic("builtin imaginary numbers do not work with V")
		case token.CHAR:
			c.write("`" + expr.Value[1:len(expr.Value)-1] + "`")
		case token.STRING:
			c.write("'" + expr.Value[1:len(expr.Value)-1] + "'")
		}
	default:
		panic(fmt.Sprintf("unknown expression: %T", _expr))
	}
}

func (c *convert) map_fn_expr(_expr ast.Expr) {
	switch expr := _expr.(type) {
	case *ast.SelectorExpr:
		switch x := expr.X.(type) {
		case *ast.Ident:
			switch x.Name {
			case "fmt":
				switch expr.Sel.Name {
				case "Println":
					c.write("println")
				}
			}
		default:
			c.expr(_expr)
		}
	default:
		c.expr(_expr)
	}
}
