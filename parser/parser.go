package parser

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"

	"github.com/json_parser/ast"
)

type Parser struct {
	Input string
	Pos   int
}

// creates a new parser instance with input as json_string ans pos as 0
func NewParser(input string) *Parser {
	return &Parser{Input: input, Pos: 0}
}

func (p *Parser) SkipWhiteSpace() {
	for p.Pos < len(p.Input) && unicode.IsSpace(rune(p.Input[p.Pos])) {
		p.Pos++
	}
}

func (p *Parser) Parse() (ast.ASTNode, error) {
	p.SkipWhiteSpace()
	if p.Pos >= len(p.Input) {
		return ast.ASTNode{}, errors.New("unexpected end of input")
	}

	switch p.Input[p.Pos] {
	case '{':
		return p.ParseObject()
	case '[':
		return p.ParseArray()
	case '"':
		return p.ParseString()
	case 't', 'f':
		return p.ParseBoolean()
	case 'n':
		return p.ParseNull()
	default:
		if unicode.IsDigit(rune(p.Input[p.Pos])) || p.Input[p.Pos] == '-' {
			return p.ParseNumber()
		}
	}

	return ast.ASTNode{}, errors.New("unexpected character")
}

func (p *Parser) ParseObject() (ast.ASTNode, error) {
	obj := ast.ObjectNode{Pairs: make(map[string]ast.ASTNode)}
	p.Pos++

	for p.Pos < len(p.Input) {

		p.SkipWhiteSpace()

		if p.Input[p.Pos] == '}' {
			p.Pos++
			return ast.ASTNode{Type: ast.NodeObject, Value: obj}, nil
		}

		key, err := p.ParseString()

		if err != nil {
			return ast.ASTNode{}, err
		}

		p.SkipWhiteSpace()

		if p.Pos >= len(p.Input) || p.Input[p.Pos] != ':' {
			return ast.ASTNode{}, errors.New("expected ':'")
		}

		p.Pos++

		value, err := p.Parse()

		if err != nil {
			return ast.ASTNode{}, err
		}

		obj.Pairs[key.Value.(string)] = value

		p.SkipWhiteSpace()

		if p.Pos >= len(p.Input) {
			return ast.ASTNode{}, errors.New("unexpected end of input")
		}

		if p.Input[p.Pos] == ',' {
			p.Pos++
		} else if p.Input[p.Pos] != '}' {
			fmt.Println(p.Input[p.Pos:len(p.Input)])
			fmt.Println("position", p.Pos)
			return ast.ASTNode{}, errors.New("expected ',' or '}'")
		}
	}

	return ast.ASTNode{}, errors.New("unterminated object")
}

func (p *Parser) ParseArray() (ast.ASTNode, error) {
	arr := ast.ArrayNode{Elements: make([]ast.ASTNode, 0)}
	p.Pos++

	for p.Pos < len(p.Input) {
		p.SkipWhiteSpace()
		if p.Input[p.Pos] == ']' {
			p.Pos++
			return ast.ASTNode{Type: ast.NodeArray, Value: arr}, nil
		}

		value, err := p.Parse()

		if err != nil {
			return ast.ASTNode{}, err
		}

		arr.Elements = append(arr.Elements, value)

		p.SkipWhiteSpace()

		if p.Pos >= len(p.Input) {
			return ast.ASTNode{}, errors.New("unexpected end of input")
		}

		if p.Input[p.Pos] == ',' {
			p.Pos++
		} else if p.Input[p.Pos] != ']' {
			return ast.ASTNode{}, errors.New("expected ',' or ']'")
		}
	}

	return ast.ASTNode{}, errors.New("unterminated array")
}
func (p *Parser) ParseString() (ast.ASTNode, error) {
	if p.Input[p.Pos] != '"' {
		return ast.ASTNode{}, errors.New("expected '\"'")
	}
	p.Pos++

	start := p.Pos

	for p.Pos < len(p.Input) {
		if p.Input[p.Pos] == '"' && p.Input[p.Pos-1] != '\\' {
			result := p.Input[start:p.Pos]
			p.Pos++
			return ast.ASTNode{Type: ast.NodeString, Value: result}, nil
		}
		p.Pos++
	}

	return ast.ASTNode{}, errors.New("unterminated string")
}

func (p *Parser) ParseNumber() (ast.ASTNode, error) {
	start := p.Pos
	for p.Pos < len(p.Input) && (unicode.IsDigit(rune(p.Input[p.Pos])) || p.Input[p.Pos] == '.' || p.Input[p.Pos] == '-' || p.Input[p.Pos] == 'e' || p.Input[p.Pos] == 'E') {
		p.Pos++
	}
	num, err := strconv.ParseFloat(p.Input[start:p.Pos], 64)
	if err != nil {
		return ast.ASTNode{}, err
	}

	return ast.ASTNode{Type: ast.NodeNumber, Value: num}, nil
}
func (p *Parser) ParseBoolean() (ast.ASTNode, error) {
	if p.Pos+4 <= len(p.Input) && p.Input[p.Pos:p.Pos+4] == "true" {
		p.Pos += 4
		return ast.ASTNode{Type: ast.NodeBoolean, Value: true}, nil
	}
	if p.Pos+5 <= len(p.Input) && p.Input[p.Pos:p.Pos+5] == "false" {
		p.Pos += 5
		return ast.ASTNode{Type: ast.NodeBoolean, Value: false}, nil
	}

	return ast.ASTNode{}, errors.New("invalid boolean")
}
func (p *Parser) ParseNull() (ast.ASTNode, error) {
	if p.Pos+4 <= len(p.Input) && p.Input[p.Pos:p.Pos+4] == "null" {
		p.Pos += 4
		return ast.ASTNode{Type: ast.NodeNull, Value: nil}, nil
	}
	return ast.ASTNode{}, errors.New("invalid null")
}
