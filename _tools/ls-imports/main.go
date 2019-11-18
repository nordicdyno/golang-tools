package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

var (
	underscoresOnly = flag.Bool("u", false, "install only underscore imports")
	parseFile       = flag.String("f", "tools.go", "path to file with imports")
)

func main() {
	flag.Parse()
	// fmt.Println("os.Args:", os.Args[0])
	f, err := parser.ParseFile(token.NewFileSet(), *parseFile, nil, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	v := &visitor{}
	ast.Walk(v, f)
	for _, imp := range v.imports {
		if *underscoresOnly && imp.Name != "_" {
			continue
		}
		fmt.Println(imp.Value)
	}
}

type importItem struct {
	Name  string
	Value string
}

type visitor struct {
	imports []importItem
}

func (v *visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}

	if imp, ok := n.(*ast.ImportSpec); ok {
		v.imports = append(v.imports, importItem{
			Name:  imp.Name.String(),
			Value: imp.Path.Value,
		})
	}
	return v
}
