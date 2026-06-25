package lexer

import (
    "fmt"
    "github.com/Sameetpatro/odlang/token"
)

type Lexer struct {
    input    string
    position int  
    readpos  int  
    ch       byte 
}

func New(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar()
    return l
}

func (l *Lexer) readChar() {
    if l.readpos >= len(l.input) {
        l.ch = 0 // 0 means EOF, same as ASCII NUL
    } else {
        l.ch = l.input[l.readpos]
    }
    l.position = l.readpos
    l.readpos++
}

//peekchar next position check kariki kuhe agaku kn achi 
func (l *Lexer) peekChar() byte {
    if l.readpos >= len(l.input) {
        return 0
    }
    return l.input[l.readpos]
}

func (l *Lexer) skipWhitespace() {
    for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
        l.readChar()
    }
}

//aeita sabu valid identifiers check kariki kuhe and also includes _
func isLetter(ch byte) bool {
    return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

// readIdentifier slices directly from the input string — no allocations.
// l.input[start:l.position] gives us the identifier without copying bytes.
func (l *Lexer) readIdentifier() string {
    start := l.position
    for isLetter(l.ch) {
        l.readChar()
    }
    return l.input[start:l.position]
}

//aeita string read kare until " "" " or EOF
//jo string ase seita is without quotes
func (l *Lexer) readString() string {
    l.readChar() // skip opening "
    start := l.position
    for l.ch != '"' && l.ch != 0 {
        l.readChar()
    }
    str := l.input[start:l.position]
    l.readChar() // skip closing "
    return str
}

// NextToken is the heart of the lexer.
// It skips whitespace, looks at the current character, and returns the next token.
// IMPORTANT: for single-char tokens, we build the token THEN call readChar().
// For multi-char tokens, readIdentifier/readString already advance internally.
func (l *Lexer) NextToken() token.Token {
    l.skipWhitespace()

    var tok token.Token

    switch l.ch {
    case '(':
        tok = token.Token{Type: token.LPAREN, Literal: "("}
        l.readChar()
    case ')':
        tok = token.Token{Type: token.RPAREN, Literal: ")"}
        l.readChar()
    case '"':
        tok = token.Token{Type: token.STRING, Literal: l.readString()}
    case 0:
        tok = token.Token{Type: token.EOF, Literal: ""}
    default:
        if isLetter(l.ch) {
            literal := l.readIdentifier()
            // LookupIdent distinguishes keywords like 'lekha' from plain identifiers
            tok = token.Token{Type: token.LookupIdent(literal), Literal: literal}
        } else {
            // Anything we don't recognise becomes ILLEGAL
            tok = token.Token{Type: token.ILLEGAL, Literal: string(l.ch)}
            l.readChar()
        }
    }

    return tok
}

func (l *Lexer) Debug() {
    fmt.Printf("input=%s\n", l.input)
    fmt.Printf("position=%d\n", l.position)
    fmt.Printf("readPosition=%d\n", l.readpos)
    fmt.Printf("ch=%c\n", l.ch)
}