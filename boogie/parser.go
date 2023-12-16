package boogie

import "fmt"

type Parser struct {
	In chan Lexeme
}

func NewParser() *Parser {
	return &Parser{
		In: make(chan Lexeme),
	}
}

func (parser *Parser) Run() {
	for {
		lex := <-parser.In
		fmt.Println(lex)
	}
}
