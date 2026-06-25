package parser

import (
    "fmt"

    "github.com/Sameetpatro/odlang/ast"
    "github.com/Sameetpatro/odlang/lexer"
    "github.com/Sameetpatro/odlang/token"
)

type Parser struct {
    l       *lexer.Lexer
    curTok  token.Token
    peekTok token.Token
    errors  []string
}

// New creates a parser. We call nextToken twice to fill curTok and peekTok.
func New(l *lexer.Lexer) *Parser {
    p := &Parser{l: l}
    p.nextToken()
    p.nextToken()
    return p
}

func (p *Parser) nextToken() {
    p.curTok = p.peekTok
    p.peekTok = p.l.NextToken()
}

func (p *Parser) Errors() []string { return p.errors }

// ParseProgram loops over tokens and builds the statement list.
func (p *Parser) ParseProgram() *ast.Program {
    program := &ast.Program{}
    for p.curTok.Type != token.EOF {
        stmt := p.parseStatement()
        if stmt != nil {
            program.Statements = append(program.Statements, stmt)
        }
        p.nextToken()
    }
    return program
}

func (p *Parser) parseStatement() ast.Statement {
    switch p.curTok.Type {
    case token.LEKHA:
        return p.parseLekhaStatement()
    default:
        p.errors = append(p.errors, fmt.Sprintf("unknown token: %s", p.curTok.Literal))
        return nil
    }
}

// parseLekhaStatement expects: lekha ( "string" )
func (p *Parser) parseLekhaStatement() *ast.LekhaStatement {
    // expect '('
    p.nextToken()
    if p.curTok.Type != token.LPAREN {
        p.errors = append(p.errors, "expected '(' after lekha")
        return nil
    }
    // expect string argument
    p.nextToken()
    if p.curTok.Type != token.STRING {
        p.errors = append(p.errors, "expected string argument in lekha()")
        return nil
    }
    value := p.curTok.Literal
    // expect ')'
    p.nextToken()
    if p.curTok.Type != token.RPAREN {
        p.errors = append(p.errors, "expected ')' after lekha argument")
        return nil
    }
    return &ast.LekhaStatement{Value: value}
}