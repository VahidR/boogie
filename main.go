package main

/**
 * Created by vahid (@vahid_r)
 * "Good Thoughts, Good Words, Good Deeds. --Zarathustra"
 */

import (
	"bufio"
	"fmt"
	"github.com/VahidR/boogie/boogie"
	"os"
)

func main() {
	lexerInChannel := make(chan string)
	lexer := boogie.NewLexer(lexerInChannel)
	go lexer.Run()
	// an endless loop to interpret whatever users enter. A kind of REPL!
	for {
		// this is how we read from STDIN in golang
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">>> ")                    // Python, is that you?
		text, err := reader.ReadString('\n') // this is single quote as we only want to read a single char (\n)
		if err != nil {
			panic(err)
		}
		lexer.In <- text
	}
}
