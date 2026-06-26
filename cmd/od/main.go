package main

import (
    "fmt"
    "os"

    "github.com/Sameetpatro/odlang/interpreter"
    "github.com/Sameetpatro/odlang/lexer"
    "github.com/Sameetpatro/odlang/parser"
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
    p := parser.New(l)
    program := p.ParseProgram()

    // Print parser errors and exit if any
    if len(p.Errors()) > 0 {
        for _, e := range p.Errors() {
            fmt.Println("parse error:", e)
        }
        os.Exit(1)
    }

    interpreter.Eval(program)
}