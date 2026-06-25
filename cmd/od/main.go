package main

import (
    "fmt"
    "os"
    "github.com/Sameetpatro/odlang/lexer"
    "github.com/Sameetpatro/odlang/token"
)

func main() {
    if len(os.Args) < 3 || os.Args[1] != "run" {
        fmt.Println("Usage: od run <file.od>")
        return
    }

    src, err := os.ReadFile(os.Args[2])
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    l := lexer.New(string(src))

    for {
        tok := l.NextToken()
        fmt.Printf("%-10s %q\n", tok.Type, tok.Literal)
        if tok.Type == token.EOF {
            break
        }
    }
}