package boogie

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
		// lex := <-parser.In
	}
}
