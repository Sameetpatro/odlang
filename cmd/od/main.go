package main

import (
	"fmt"
	"os"

	"github.com/Sameetpatro/odlang/lexer"
	"github.com/Sameetpatro/odlang/token"
)

// main starts the od command and runs Phase A token printing.
// Example: od run hello.od reads the file and prints every token
func main() {
	if len(os.Args) < 3 || os.Args[1] != "run" {
		fmt.Println("Usage: od run <file.od>")
		return
	}

	source, err := os.ReadFile(os.Args[2])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lexerInstance := lexer.New(string(source))
	for {
		currentToken := lexerInstance.NextToken()
		printToken(currentToken)
		if currentToken.Type == token.EOF {
			break
		}
	}
}

// printToken shows one token in the verification format.
// Example: LEKHA token prints as: LEKHA       "lekha"
func printToken(currentToken token.Token) {
	fmt.Printf("%-12s %q\n", currentToken.Type, currentToken.Literal)
}
