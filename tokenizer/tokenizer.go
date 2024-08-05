package tokenizer

import (
	"fmt"
	"unicode"
)

type TokenType int

const (
	TokenNumber TokenType = iota
	TokenPlus
	TokenMinus
	TokenMultiply
	TokenDivide
	TokenLeftParen
	TokenRightParen
)

type Token struct {
	Type  TokenType
	Value string
}

func Tokenize(expression string) ([]Token, error) {
	var tokens []Token
	for i := 0; i < len(expression); i++ {
		char := expression[i]
		switch {
		case unicode.IsDigit(rune(char)):
			dotFound := false
			number := string(char)
			j := i + 1
			for ; j < len(expression); j++ {
				if unicode.IsDigit(rune(expression[j])) || (expression[j] == '.' && !dotFound) {
					number += string(expression[j])
					if expression[j] == '.' {
						dotFound = true
					}
				} else {
					break
				}
			}
			tokens = append(tokens, Token{Type: TokenNumber, Value: number})

		case char == '+':
			tokens = append(tokens, Token{Type: TokenPlus, Value: string(char)})
		case char == '-':
			tokens = append(tokens, Token{Type: TokenMinus, Value: string(char)})
		case char == '*':
			tokens = append(tokens, Token{Type: TokenMultiply, Value: string(char)})
		case char == '/':
			tokens = append(tokens, Token{Type: TokenDivide, Value: string(char)})
		case char == '(':
			tokens = append(tokens, Token{Type: TokenLeftParen, Value: string(char)})
		case char == ')':
			tokens = append(tokens, Token{Type: TokenRightParen, Value: string(char)})
		case unicode.IsSpace(rune(char)):
			continue
		default:
			return nil, fmt.Errorf("unexpected character: %v", string(char))
		}
	}
	return tokens, nil
}
