# mathsolve-go

## Description

This is a simple math solver written in Go. It can solve simple math problems like addition, subtraction, multiplication, and division from plain text.

## Usage

```go
  go mod tidy
  go run main.go -e "your-expression-here"
```

## Example

```go
  go run main.go -e "2 + 2"
```

## How it works

the program is created as if it is a compiler . it envolves three steps :

### 1. Lexical Analysis

This is the first step in the compiler process. It takes the input string and converts it into a sequence of tokens. In this program, the input is a simple math expression like 2 + 2, and the output will be a sequence of tokens such as NUMBER(2) PLUS NUMBER(2). Mathematically, it can be described as a finite state automaton processing a regular grammar.

#### Tokens

- TokenNumber : represents a number like 2 , 12.4 , 0.5
- TokenPlus : represents the plus operator
- TokenMinus : represents the minus operator
- TokenMultiply : represents the multiply operator
- TokenDivide : represents the divide operator
- TokenLeftParen : represents the left parenthesis
- TokenRightParen : represents the right parenthesis

#### Grammar

- Expression : Term | Expression PLUS Term | Expression MINUS Term
- Term : Factor | Term MULTIPLY Factor | Term DIVIDE Factor
- Factor : NUMBER | LEFT_PAREN Expression RIGHT_PAREN

### 2. Parsing

The second step is parsing. It takes the sequence of tokens produced by the lexical analysis and converts it into an abstract syntax tree (AST). In this program, the token sequence NUMBER(2) PLUS NUMBER(2) will be converted into an AST like ADD(NUMBER(2), NUMBER(2)). This can be viewed as a pushdown automaton processing a context-free grammar. The AST representation will look like this:

```
  ADD
  / \
  2   2
```

#### Abstract Syntax Tree

For the generation of the abstract syntax tree , we should take in consideration the precedence of the operators , for example the multiplication and division operators should be evaluated before the addition and subtraction operators , so the abstract syntax tree should look like this :

```
  ADD
  / \
  2   MULTIPLY
      / \
      2   2
```

For the algorithm of the generation of the abstract syntax tree , we can use the [Recursive Descent Parsing](https://en.wikipedia.org/wiki/Recursive_descent_parser#:~:text=In%20computer%20science%2C%20a%20recursive,the%20nonterminals%20of%20the%20grammar.) algorithm , which is a top-down parsing algorithm that constructs the parse tree from the top and the input is read from left to right.

### 3. Evaluation

The final step is evaluation. It takes the AST and evaluates it. In this program, the AST ADD(NUMBER(2), NUMBER(2)) will be evaluated to 4. This step can be reviewed as a Turing machine processing a context-sensitive grammar.

#### Evaluation

For the evaluation of the abstract syntax tree , we can use the recursive evaluation algorithm , which is a top-down evaluation algorithm that evaluates the parse tree from the top and the input is read from left to right.
For example , the evaluation of the abstract syntax tree ADD(NUMBER(2), MULTIPLY(NUMBER(2), NUMBER(2))) will be :

```
  ADD
  / \
  2   MULTIPLY
      / \
      2   2
```

The evaluation of the abstract syntax tree will be :

1. Evaluate the left child of the ADD node which is NUMBER(2). The result will be 2.
2. Evaluate the right child of the ADD node which is MULTIPLY(NUMBER(2), NUMBER(2)).
3. Evaluate the left child of the MULTIPLY node which is NUMBER(2). The result will be 2.
4. Evaluate the right child of the MULTIPLY node which is NUMBER(2). The result will be 2.
5. Multiply the results of the left and right children of the MULTIPLY node which will be 4.
6. Add the results of the left and right children of the ADD node which will be 6.
