package parser

import "github.com/edaywalid/mathsolve-go/tokenizer"

type NodeType int

const (
	NodeNumber NodeType = iota
	NodeOperator
)

type Node struct {
	Type  NodeType
	Value string
	Left  *Node
	Right *Node
}

func NewNode(t NodeType, v string) *Node {
	return &Node{Type: t, Value: v}
}

type Parser struct {
	tokens  []tokenizer.Token
	current int
}

func NewParser(tokens []tokenizer.Token) *Parser {
	return &Parser{tokens: tokens, current: 0}
}

func (p *Parser) parse() *Node {
	return p.parseExpression()
}

func (p *Parser) parseExpression() *Node {
	node := p.parseTerm()

	for p.current < len(p.tokens) && (p.tokens[p.current].Type == tokenizer.TokenPlus || p.tokens[p.current].Type == tokenizer.TokenMinus) {
		token := p.tokens[p.current]
		p.current++
		node = &Node{
			Type:  NodeOperator,
			Value: token.Value,
			Left:  node,
			Right: p.parseTerm(),
		}
	}

	return node
}

func (p *Parser) parseTerm() *Node {
	node := p.parseFactor()

	for p.current < len(p.tokens) && (p.tokens[p.current].Type == tokenizer.TokenMultiply || p.tokens[p.current].Type == tokenizer.TokenDivide) {
		token := p.tokens[p.current]
		p.current++
		node = &Node{
			Type:  NodeOperator,
			Value: token.Value,
			Left:  node,
			Right: p.parseFactor(),
		}
	}

	return node
}

func (p *Parser) parseFactor() *Node {
	token := p.tokens[p.current]

	if token.Type == tokenizer.TokenNumber {
		p.current++
		return &Node{
			Type:  NodeNumber,
			Value: token.Value,
		}
	}

	if token.Type == tokenizer.TokenLeftParen {
		p.current++
		node := p.parseExpression()
		if p.tokens[p.current].Type != tokenizer.TokenRightParen {
			panic("missing closing parenthesis")
		}
		p.current++
		return node
	}

	panic("unexpected token")
}

func Parse(tokens []tokenizer.Token) *Node {
	parser := NewParser(tokens)
	return parser.parse()
}
