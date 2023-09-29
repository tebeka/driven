package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("relu.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", string(data), 0)
	if err != nil {
		log.Fatal(err)
	}

	for _, d := range f.Decls {
		if fn, ok := d.(*ast.FuncDecl); ok {
			fmt.Println(fn.Name)
			for _, typ := range fn.Type.TypeParams.List {
				printType(typ.Type)
			}
		}
	}
}

func printType(typ ast.Expr) {
	switch v := typ.(type) {
	case *ast.BinaryExpr: // int | float64
		printType(v.X)
		printType(v.Y)
	case *ast.UnaryExpr: // ~int
		printType(v.X)
	default:
		fmt.Println(typ)
	}
}
