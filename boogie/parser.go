package boogie

import "github.com/davecgh/go-spew/spew"

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
		spew.Dump(lex)
	}
}
