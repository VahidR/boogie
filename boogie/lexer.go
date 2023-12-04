package boogie

/**
 * Created by vahid (@vahid_r)
 * "Good Thoughts, Good Words, Good Deeds. --Zarathustra"
 */

import (
	"strings"
)

// A State-Machine Enum
type ContextState int

const (
	START ContextState = iota
	KEYWORD
)

type Lexer struct {
	Input          chan string // the channel that we use to read from STDIN
	currentChar    string
	currentContext ContextState // the current state of the lexer
	keywords       []string
	buffer         string // the buffer that we use to store the current token
}

// NewLexer Constructor for the Lexer
func NewLexer(input chan string) *Lexer {
	return &Lexer{
		Input:          input,
		currentContext: START,
		keywords:       []string{"QUIT"}, // the keywords that we want to recognize
	}
}

// Run runs the lexer
func (lexer *Lexer) Run() {
	for {
		lineOfCode := <-lexer.Input
		err := lexer.process(lineOfCode)
		if err != nil {
			panic(err)
		}
	}
}

func (lexer *Lexer) process(lineOfCode string) error {
	for _, r := range lineOfCode {
		lexer.currentChar = string(r) // convert it to string
		switch lexer.currentContext {
		case START:
			lexer.buffer += lexer.currentChar
			if lexer.isStringInSlice(lexer.buffer, lexer.keywords) {
				lexer.currentContext = KEYWORD
			}
		case KEYWORD:
			// do something with the keyword
		}
	}
	return nil
}

func (lexer *Lexer) isStringInSlice(str string, stringList []string) bool {
	for _, value := range stringList {
		if value == strings.ToUpper(str) {
			return true
		}
	}
	return false
}
