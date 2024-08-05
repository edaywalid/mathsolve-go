package evaluator

import (
	"strconv"

	"github.com/edaywalid/mathsolve-go/parser"
)

type evaluator struct {
	node *parser.Node
}

func NewEvaluator(node *parser.Node) *evaluator {
	return &evaluator{node: node}
}

func (e *evaluator) Evaluate() float64 {
	return e.evaluateNode(e.node)
}

func (e *evaluator) evaluateNode(node *parser.Node) float64 {
	if node == nil {
		return 0
	}

	if node.Type == parser.NodeNumber {
		return e.evaluateNumber(node)
	}

	return e.evaluateOperator(node)
}

func (e *evaluator) evaluateNumber(node *parser.Node) float64 {
	i, err := strconv.Atoi(node.Value)
	if err != nil {
		panic(err)
	}
	return float64(i)
}

func (e *evaluator) evaluateOperator(node *parser.Node) float64 {
	left := e.evaluateNode(node.Left)
	right := e.evaluateNode(node.Right)

	switch node.Value {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	}

	return 0
}
