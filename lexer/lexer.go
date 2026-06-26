package lexer

import (
	"github.com/Sameetpatro/odlang/token"
)

// Lexer reads the source code one character at a time.
// It turns raw text into tokens that the parser can understand.
// Example: "lekha("hi")" becomes [LEKHA, LPAREN, STRING("hi"), RPAREN]
type Lexer struct {
	input        string
	position     int
	readPosition int
	currentChar  byte
}

// New creates a lexer ready to read the given source code.
// Example: New("lekha(1)") returns a lexer at the first character 'l'
func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

// readChar moves to the next character in the input string.
// Example: after readChar on "abc", currentChar becomes 'b'
func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.currentChar = 0
	} else {
		lexer.currentChar = lexer.input[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition++
}

// peekChar looks at the next character without moving forward.
// Example: if input is "10", currentChar is '1' and peekChar returns '0'
func (lexer *Lexer) peekChar() byte {
	if lexer.readPosition >= len(lexer.input) {
		return 0
	}
	return lexer.input[lexer.readPosition]
}

// skipWhitespace moves past spaces, tabs, and newlines.
// Example: "  lekha" becomes "lekha" after skipWhitespace
func (lexer *Lexer) skipWhitespace() {
	for lexer.currentChar == ' ' || lexer.currentChar == '\t' ||
		lexer.currentChar == '\n' || lexer.currentChar == '\r' {
		lexer.readChar()
	}
}

// readIdentifier reads a word made of letters and underscores.
// Example: "aarambha" in karya aarambha() is read as one identifier
func (lexer *Lexer) readIdentifier() string {
	start := lexer.position
	for isLetter(lexer.currentChar) {
		lexer.readChar()
	}
	return lexer.input[start:lexer.position]
}

// readNumber reads a whole number without a decimal point.
// Example: "10" in x = 10 | is read as the string "10"
func (lexer *Lexer) readNumber() string {
	start := lexer.position
	for isDigit(lexer.currentChar) {
		lexer.readChar()
	}
	return lexer.input[start:lexer.position]
}

// readFloat reads a number that has a dot in the middle.
// Example: "3.14" in const PI = 3.14 is read as the string "3.14"
func (lexer *Lexer) readFloat() string {
	start := lexer.position
	for isDigit(lexer.currentChar) || lexer.currentChar == '.' {
		lexer.readChar()
	}
	return lexer.input[start:lexer.position]
}

// readString reads text inside double quotes.
// Example: "hello" in lekha("hello") returns hello without the quotes
func (lexer *Lexer) readString() string {
	lexer.readChar()
	start := lexer.position
	for lexer.currentChar != '"' && lexer.currentChar != 0 {
		lexer.readChar()
	}
	literal := lexer.input[start:lexer.position]
	lexer.readChar()
	return literal
}

// readChar_literal reads one character inside single quotes.
// Example: 'a' in akshara c = 'a' returns the character a without quotes
func (lexer *Lexer) readChar_literal() string {
	lexer.readChar()
	start := lexer.position
	for lexer.currentChar != '\'' && lexer.currentChar != 0 {
		lexer.readChar()
	}
	literal := lexer.input[start:lexer.position]
	lexer.readChar()
	return literal
}

// isLetter checks if a character can start or continue an identifier.
// Example: isLetter('a') is true, isLetter('5') is false
func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

// isDigit checks if a character is a number digit.
// Example: isDigit('7') is true, isDigit('x') is false
func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

// NextToken returns the next token from the source code.
// Example: "lekha(x) |" returns LEKHA, then LPAREN, then IDENT "x", etc.
func (lexer *Lexer) NextToken() token.Token {
	lexer.skipWhitespace()

	var tok token.Token

	switch lexer.currentChar {
	case '(':
		tok = newToken(token.LPAREN, lexer.currentChar)
	case ')':
		tok = newToken(token.RPAREN, lexer.currentChar)
	case '{':
		tok = newToken(token.LBRACE, lexer.currentChar)
	case '}':
		tok = newToken(token.RBRACE, lexer.currentChar)
	case '[':
		tok = newToken(token.LBRACK, lexer.currentChar)
	case ']':
		tok = newToken(token.RBRACK, lexer.currentChar)
	case '|':
		tok = newToken(token.PIPE, lexer.currentChar)
	case '+':
		if lexer.peekChar() == '+' {
			ch := lexer.currentChar
			lexer.readChar()
			tok = token.Token{Type: token.INCREMENT, Literal: string(ch) + string(lexer.currentChar)}
		} else {
			tok = newToken(token.PLUS, lexer.currentChar)
		}
	case '-':
		if lexer.peekChar() == '>' {
			ch := lexer.currentChar
			lexer.readChar()
			tok = token.Token{Type: token.ARROW, Literal: string(ch) + string(lexer.currentChar)}
		} else if lexer.peekChar() == '-' {
			ch := lexer.currentChar
			lexer.readChar()
			tok = token.Token{Type: token.DECREMENT, Literal: string(ch) + string(lexer.currentChar)}
		} else {
			tok = newToken(token.MINUS, lexer.currentChar)
		}
	case '/':
		tok = newToken(token.SLASH, lexer.currentChar)
	case ',':
		tok = newToken(token.COMMA, lexer.currentChar)
	case '"':
		tok = token.Token{Type: token.STRING, Literal: lexer.readString()}
	case '\'':
		tok = token.Token{Type: token.CHAR, Literal: lexer.readChar_literal()}
	case '!':
		if lexer.peekChar() == '=' {
			ch := lexer.currentChar
			lexer.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(lexer.currentChar)}
		} else {
			tok = newToken(token.BANG, lexer.currentChar)
		}
	case '>':
		if lexer.peekChar() == '=' {
			ch := lexer.currentChar
			lexer.readChar()
			tok = token.Token{Type: token.GTE, Literal: string(ch) + string(lexer.currentChar)}
		} else if lexer.peekChar() == '>' {
			ch := lexer.currentChar
			lexer.readChar()
			tok = token.Token{Type: token.RSHIFT, Literal: string(ch) + string(lexer.currentChar)}
		} else {
			tok = newToken(token.GT, lexer.currentChar)
		}
	case '<':
		if lexer.peekChar() == '=' {
			ch := lexer.currentChar
			lexer.readChar()
			tok = token.Token{Type: token.LTE, Literal: string(ch) + string(lexer.currentChar)}
		} else {
			tok = newToken(token.LT, lexer.currentChar)
		}
	case '=':
		if lexer.peekChar() == '=' {
			ch := lexer.currentChar
			lexer.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(lexer.currentChar)}
		} else {
			tok = newToken(token.ASSIGN, lexer.currentChar)
		}
	case '*':
		if lexer.peekChar() == '*' {
			ch := lexer.currentChar
			lexer.readChar()
			tok = token.Token{Type: token.EXP, Literal: string(ch) + string(lexer.currentChar)}
		} else {
			tok = newToken(token.STAR, lexer.currentChar)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(lexer.currentChar) {
			literal := lexer.readIdentifier()
			tok.Type = token.LookupIdent(literal)
			tok.Literal = literal
			return tok
		} else if isDigit(lexer.currentChar) {
			if lexer.peekChar() == '.' {
				literal := lexer.readFloat()
				tok = token.Token{Type: token.FLOAT, Literal: literal}
			} else {
				literal := lexer.readNumber()
				tok = token.Token{Type: token.INT, Literal: literal}
			}
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lexer.currentChar)
		}
	}

	lexer.readChar()
	return tok
}

// newToken builds a token from a type and one character.
// Example: newToken(token.PLUS, '+') gives {Type: "+", Literal: "+"}
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
