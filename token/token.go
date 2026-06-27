package token

// TokenType tells us what kind of token we found in the source code.
// Example: the word "lekha" has type LEKHA, the number 10 has type INT
type TokenType string

// Token is one piece of the source code after the lexer reads it.
// Example: lekha(x) gives tokens like {Type: LEKHA, Literal: "lekha"}
type Token struct {
	Type    TokenType
	Literal string
}

const (
	// ILLEGAL means we found a character we do not understand.
	// Example: @ in "x = @5" becomes ILLEGAL
	ILLEGAL = "ILLEGAL"

	// EOF means we reached the end of the source file.
	// Example: after the last token, we return EOF with Literal ""
	EOF = "EOF"

	// IDENT is a name like a variable or function name.
	// Example: "aarambha" in karya aarambha() is IDENT
	IDENT = "IDENT"

	// STRING is text inside double quotes.
	// Example: "hello" in lekha("hello") is STRING
	STRING = "STRING"

	// INT is a whole number without a dot.
	// Example: 10 in sankhya x = 10 is INT
	INT = "INT"

	// FLOAT is a number with a dot.
	// Example: 3.14 in const PI = 3.14 is FLOAT
	FLOAT = "FLOAT"

	// CHAR is one character inside single quotes.
	// Example: 'a' in akshara c = 'a' is CHAR
	CHAR = "CHAR"

	// LPAREN is the left round bracket (
	// Example: lekha( uses LPAREN
	LPAREN = "LPAREN"

	// RPAREN is the right round bracket )
	// Example: lekha(x) uses RPAREN
	RPAREN = "RPAREN"

	// LBRACE is the left curly bracket {
	// Example: { after karya aarambha() { opens a block
	LBRACE = "LBRACE"

	// RBRACE is the right curly bracket }
	// Example: } at the end of a block is RBRACE
	RBRACE = "RBRACE"

	// LBRACK is the left square bracket [
	// Example: arr[0] uses LBRACK
	LBRACK = "LBRACK"

	// RBRACK is the right square bracket ]
	// Example: arr[0] uses RBRACK
	RBRACK = "RBRACK"

	// SEMI is the statement terminator ;
	// Example: lekha(x) ; ends the line with SEMI
	SEMI = "SEMI"

	// BANG is the not operator !
	// Example: !han means not true
	BANG = "BANG"

	// GT is the greater than operator >
	// Example: x > 10 uses GT
	GT = "GT"

	// LT is the less than operator <
	// Example: x < 10 uses LT
	LT = "LT"

	// ASSIGN is the equals sign for giving a value
	// Example: x = 10 uses ASSIGN
	ASSIGN = "ASSIGN"

	// PLUS is the add operator +
	// Example: 1 + 2 uses PLUS
	PLUS = "PLUS"

	// MINUS is the subtract operator -
	// Example: x - 1 uses MINUS
	MINUS = "MINUS"

	// STAR is the multiply operator *
	// Example: 2 * 3 uses STAR
	STAR = "STAR"

	// SLASH is the divide operator /
	// Example: 10 / 2 uses SLASH
	SLASH = "SLASH"

	// EQ is the equal compare operator ==
	// Example: x == 10 uses EQ
	EQ = "EQ"

	// NOT_EQ is the not equal operator !=
	// Example: x != 0 uses NOT_EQ
	NOT_EQ = "NOT_EQ"

	// GTE is the greater or equal operator >=
	// Example: x >= 10 uses GTE
	GTE = "GTE"

	// LTE is the less or equal operator <=
	// Example: x <= 10 uses LTE
	LTE = "LTE"

	// EXP is the power operator **
	// Example: 2**3 means 2 to the power 3
	EXP = "EXP"

	// RSHIFT is the input split operator >>
	// Example: dia(p >> q) uses RSHIFT between variables
	RSHIFT = "RSHIFT"

	// COMMA separates items in a list.
	// Example: misana(a, b) uses COMMA between arguments
	COMMA = "COMMA"

	// ARROW is the range operator used in ghura loops.
	// Example: ghura sankhya i = 0 -> 5 uses ARROW between start and end
	ARROW = "ARROW"

	// INCREMENT adds one to a loop variable after each step.
	// Example: i++ after ghura ... ; i++ { increases i by 1
	INCREMENT = "INCREMENT"

	// DECREMENT subtracts one from a variable after each step.
	// Example: i-- lowers i by 1 each time
	DECREMENT = "DECREMENT"

	// LEKHA is the print keyword
	// Example: lekha("hi") prints hi
	LEKHA = "LEKHA"

	// DIA is the input keyword
	// Example: dia(p >> q) reads input into p and q
	DIA = "DIA"

	// JADI is the if keyword
	// Example: jadi x > 10 { starts an if block
	JADI = "JADI"

	// NAHELE is the else keyword
	// Example: } nahele { starts the else block
	NAHELE = "NAHELE"

	// GHURA is the for loop keyword
	// Example: ghura sankhya i = 0 -> 10 starts a for loop
	GHURA = "GHURA"

	// JETEBELEJAIN is the while loop keyword
	// Example: jetebeleJain x > 0 { runs while x is greater than 0
	JETEBELEJAIN = "JETEBELEJAIN"

	// DEIDE is the return keyword
	// Example: deide (a, b) returns two values
	DEIDE = "DEIDE"

	// BAHARIPADE is the break keyword
	// Example: baharipade stops a loop early
	BAHARIPADE = "BAHARIPADE"

	// CHADIDE is the continue keyword
	// Example: chadide skips to the next loop step
	CHADIDE = "CHADIDE"

	// KARYA is the function keyword
	// Example: karya misana(...) { defines a function
	KARYA = "KARYA"

	// SRENI is the class keyword
	// Example: sreni Point { defines a class
	SRENI = "SRENI"

	// NUA is the var keyword for new variables
	// Example: nua z = 99 declares a new variable
	NUA = "NUA"

	// CONST is the constant keyword
	// Example: const PI = 3.14 makes a value that cannot change
	CONST = "CONST"

	// ANAA is the import keyword
	// Example: anaa fmt imports the fmt module
	ANAA = "ANAA"

	// CHESTA is the try keyword
	// Example: chesta { runs code that might fail
	CHESTA = "CHESTA"

	// DHARE is the catch keyword
	// Example: } dhare { handles errors from chesta
	DHARE = "DHARE"

	// KHALI is the null keyword
	// Example: nua x = khali sets x to nothing
	KHALI = "KHALI"

	// HAN is the true keyword
	// Example: satya ok = han sets ok to true
	HAN = "HAN"

	// NA is the false keyword
	// Example: satya ok = na sets ok to false
	NA = "NA"

	// AAU is the or keyword
	// Example: han aau na means true or false
	AAU = "AAU"

	// SAHITA is the and keyword
	// Example: han sahita na means true and false
	SAHITA = "SAHITA"

	// SANKHYA is the int type name
	// Example: sankhya x = 10 declares an integer
	SANKHYA = "SANKHYA"

	// SABDA is the string type name
	// Example: sabda name = "Sameet" declares a string
	SABDA = "SABDA"

	// AKSHARA is the char type name
	// Example: akshara c = 'a' declares one character
	AKSHARA = "AKSHARA"

	// DASMIC is the float type name
	// Example: dasmik pi = 3.14 declares a decimal number
	DASMIC = "DASMIC"

	// SATYA is the bool type name
	// Example: satya ok = han declares true or false
	SATYA = "SATYA"

	// KRAMA is the array type name
	// Example: krama arr(7, 0) makes an array of 7 items
	KRAMA = "KRAMA"

	// MANA is the map type name
	// Example: mana m stores key-value pairs
	MANA = "MANA"

	// THAKA is the stack type name
	// Example: thaka s is a last-in-first-out list
	THAKA = "THAKA"

	// DHADHI is the queue type name
	// Example: dhadhi q is a first-in-first-out list
	DHADHI = "DHADHI"
)

// keywords maps OdLang words to their token type.
// Example: "lekha" maps to LEKHA, "karya" maps to KARYA
var keywords = map[string]TokenType{
	"lekha":        LEKHA,
	"dia":          DIA,
	"jadi":         JADI,
	"nahele":       NAHELE,
	"ghura":        GHURA,
	"jetebeleJain": JETEBELEJAIN,
	"deide":        DEIDE,
	"baharipade":   BAHARIPADE,
	"chadide":      CHADIDE,
	"karya":        KARYA,
	"sreni":        SRENI,
	"nua":          NUA,
	"const":        CONST,
	"anaa":         ANAA,
	"chesta":       CHESTA,
	"dhare":        DHARE,
	"khali":        KHALI,
	"han":          HAN,
	"na":           NA,
	"aau":          AAU,
	"sahita":       SAHITA,
	"sankhya":      SANKHYA,
	"sabda":        SABDA,
	"akshara":      AKSHARA,
	"dasmik":       DASMIC,
	"satya":        SATYA,
	"krama":        KRAMA,
	"mana":         MANA,
	"thaka":        THAKA,
	"dhadhi":       DHADHI,
}

// LookupIdent checks if a word is a keyword or a plain name.
// Example: LookupIdent("lekha") returns LEKHA, LookupIdent("x") returns IDENT
func LookupIdent(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}
	return IDENT
}
