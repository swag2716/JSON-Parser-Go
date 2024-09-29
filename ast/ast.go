package ast

import "fmt"

type NodeType int

const (
	NodeObject NodeType = iota
	NodeArray
	NodeString
	NodeNumber
	NodeBoolean
	NodeNull
)

type ASTNode struct {
	Type  NodeType
	Value interface{}
}

type ObjectNode struct {
	Pairs map[string]ASTNode
}

type ArrayNode struct {
	Elements []ASTNode
}

func PrintAST(node ASTNode, indent string) {
	switch node.Type {
	case NodeObject:
		fmt.Println("Object:", indent)
		obj := node.Value.(ObjectNode)
		for k, v := range obj.Pairs {
			fmt.Printf("%s %s:", indent, k)
			PrintAST(v, indent+" ")
		}
	case NodeArray:
		fmt.Println(indent, "Array: ")
		arr := node.Value.(ArrayNode)
		for _, v := range arr.Elements {
			fmt.Printf("%s ", indent)
			PrintAST(v, indent+" ")
		}
	case NodeString:
		fmt.Println(indent, "String:", node.Value)
	case NodeNumber:
		fmt.Println(indent, "Number:", node.Value)
	case NodeBoolean:
		fmt.Println(indent, "Boolean:", node.Value)
	case NodeNull:
		fmt.Println(indent, "Null:", node.Value)
	}
}
