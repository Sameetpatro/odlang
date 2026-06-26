package interpreter

import (
    "fmt"

    "github.com/Sameetpatro/odlang/ast"
)

// Eval walks the AST and executes each statement.
func Eval(program *ast.Program) {
    for _, stmt := range program.Statements {
        evalStatement(stmt)
    }
}

func evalStatement(stmt ast.Statement) {
    switch s := stmt.(type) {
    // type switch: Go checks the concrete type at runtime
    case *ast.LekhaStatement:
        fmt.Println(s.Value)
    }
}