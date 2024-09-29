package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/json_parser/ast"
	"github.com/json_parser/parser"
)

func main() {
	filename := os.Args[1]
	fmt.Println("filename", filename)

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("error", err)
		return
	}

	json_string := string(content)

	fmt.Println(json_string)

	parser := parser.NewParser(json_string) //creating a new parser

	result, err := parser.Parse()

	if err != nil {
		panic(err)
	}

	ast.PrintAST(result, "")

}
