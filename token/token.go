package token

type TokenType string

type Token struct {
    Type    TokenType
    Literal string
}

const (
    ILLEGAL = "ILLEGAL"
    EOF     = "EOF"

    IDENT  = "IDENT"
    STRING = "STRING"
    INT    = "INT"

    LPAREN    = "("
    RPAREN    = ")"
    SEMICOLON = ";"

    LEKHA = "LEKHA"
)

var keywords = map[string]TokenType{
    "lekha": LEKHA,
}

// aeita check kare ki the next word gote identifier ki nahi, if identifier then aeita se identifier return kare
func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENT
}