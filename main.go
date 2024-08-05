package main

import (
	"flag"
	"fmt"

	"github.com/edaywalid/mathsolve-go/evaluator"
	"github.com/edaywalid/mathsolve-go/parser"
	"github.com/edaywalid/mathsolve-go/tokenizer"
)

func main() {
	expression := flag.String("e", "", "Mathematical expression to evaluate")

	flag.Parse()

	if *expression == "" {
		fmt.Println(" provide a mathematical expression using the -e flag")
		return
	}

	tokens, err := tokenizer.Tokenize(*expression)
	if err != nil {
		panic(err)
	}

	parser := parser.Parse(tokens)
	eval := evaluator.NewEvaluator(parser)

	res := eval.Evaluate()
	if res == float64(int64(res)) {
		fmt.Print("result is ", int64(res))
	} else {
		fmt.Print("result is ", res)
	}
}
