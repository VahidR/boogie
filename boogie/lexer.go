package boogie

import "fmt"

type Lexer struct {
	In chan string
}

// NewLexer Constructor for the Lexer
func NewLexer(in chan string) *Lexer {
	return &Lexer{
		In: in,
	}
}

func (lexer *Lexer) Run() {
	for {
		message := <-lexer.In
		fmt.Println(message)
	}
}
