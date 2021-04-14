package convert

import (
	"go/parser"
	"go/token"
	"log"
)

func convert(filename string) string {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "test.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
}
