package token

type TokenType string

type Token struct{
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT  = "IDENT"
	STRING = "STRING"

	LPAREN = "("
	RPAREN = ")"

	LEKHA = "LEKHA"
)

var keywords = map[string]TokenType{
	"lekha" : LEKHA,
}

func LookupIdent(ident string) TokenType{
	tok, ok := keywords[ident]
	if ok {
		return tok
	}
	return IDENT
}