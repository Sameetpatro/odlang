package lexer

import "fmt"

type Lexer struct{
	input string
	position int
	readpos int
	ch byte
}

func New(input string) *Lexer{
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}
func (l *Lexer) readChar() {
	if l.readpos >= len(l.input){
		l.ch = 0
	} else {
		l.ch = l.input[l.readpos]
	}
	l.position = l.readpos
	l.readpos++
}

func (l *Lexer) Debug() {
    fmt.Printf("input=%s\n", l.input)
    fmt.Printf("position=%d\n", l.position)
    fmt.Printf("readPosition=%d\n", l.readpos)
    fmt.Printf("ch=%c\n", l.ch)
}

func (l *Lexer) peekChar() byte {
	if l.readpos >= len(l.input){
		return 0
	}
	return l.input[l.readpos]
}

func (l *Lexer) skipWhitespaces(){
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}