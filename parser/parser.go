package parser

import (
	"fmt"

	"github.com/Sameetpatro/odlang/ast"
	"github.com/Sameetpatro/odlang/lexer"
	"github.com/Sameetpatro/odlang/token"
)

// Parser reads tokens and builds an abstract syntax tree.
// Example: tokens for lekha("hi") | become a LekhaStatement node
type Parser struct {
	lexerInstance *lexer.Lexer
	currentToken  token.Token
	peekToken     token.Token
	errors        []string
}

// New creates a parser and loads the first two tokens for lookahead.
// Example: New(lexer) returns a parser ready to read from the lexer
func New(lexerInstance *lexer.Lexer) *Parser {
	parser := &Parser{lexerInstance: lexerInstance}
	parser.nextToken()
	parser.nextToken()
	return parser
}

// nextToken moves currentToken forward and reads one new token from the lexer.
// Example: after nextToken, currentToken becomes what peekToken was
func (parser *Parser) nextToken() {
	parser.currentToken = parser.peekToken
	parser.peekToken = parser.lexerInstance.NextToken()
}

// Errors returns all parse error messages collected so far.
// Example: missing ")" adds a message to the errors slice
func (parser *Parser) Errors() []string {
	return parser.errors
}

// ParseProgram reads the whole file and returns the root program node.
// Example: a file with two karya functions returns Program with two statements
func (parser *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	for parser.currentToken.Type != token.EOF {
		statement := parser.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		parser.syncAfterTopLevelStatement()
	}
	return program
}

// parseStatement picks the right parser based on the current token.
// Example: currentToken LEKHA calls parseLekhaStatement
func (parser *Parser) parseStatement() ast.Statement {
	switch parser.currentToken.Type {
	case token.LEKHA:
		return parser.parseLekhaStatement()
	case token.DIA:
		return parser.parseDiaStatement()
	case token.DEIDE:
		return parser.parseDeideStatement()
	case token.JADI:
		return parser.parseIfStatement()
	case token.GHURA:
		return parser.parseGhuraStatement()
	case token.JETEBELEJAIN:
		return parser.parseJetebeleJainStatement()
	case token.KARYA:
		return parser.parseKaryaStatement()
	case token.CHESTA:
		return parser.parseChestaStatement()
	case token.BAHARIPADE:
		return parser.parseBaharipadeStatement()
	case token.CHADIDE:
		return parser.parseChadideStatement()
	case token.CONST:
		return parser.parseConstStatement()
	case token.NUA:
		return parser.parseVarStatement("", false)
	case token.SANKHYA, token.SABDA, token.AKSHARA, token.DASMIC, token.SATYA,
		token.KRAMA, token.MANA, token.THAKA, token.DHADHI:
		return parser.parseVarStatement(parser.currentToken.Literal, false)
	case token.IDENT:
		return parser.parseIdentStatement()
	case token.RBRACE:
		return nil
	default:
		parser.errors = append(parser.errors,
			fmt.Sprintf("unknown statement start: %s (%q)", parser.currentToken.Type, parser.currentToken.Literal))
		return nil
	}
}

// isTypeToken checks if a token is an OdLang type name keyword.
// Example: token.SANKHYA is a type token, token.IDENT is not
func isTypeToken(tokenType token.TokenType) bool {
	switch tokenType {
	case token.SANKHYA, token.SABDA, token.AKSHARA, token.DASMIC, token.SATYA,
		token.KRAMA, token.MANA, token.THAKA, token.DHADHI:
		return true
	default:
		return false
	}
}

// expectPeek checks the next token type and moves forward if it matches.
// Example: expectPeek(token.RPAREN) succeeds when peekToken is ")"
func (parser *Parser) expectPeek(tokenType token.TokenType) bool {
	if parser.peekToken.Type == tokenType {
		parser.nextToken()
		return true
	}
	parser.peekError(tokenType)
	return false
}

// peekError records a message when the next token is not what we expected.
// Example: wanted RPAREN but found PIPE adds an error message
func (parser *Parser) peekError(tokenType token.TokenType) {
	message := fmt.Sprintf("expected next token %s, got %s instead",
		tokenType, parser.peekToken.Type)
	parser.errors = append(parser.errors, message)
}

// isStatementStart checks if the current token can begin a new line of code.
// Example: KARYA and lekha can start statements, but PIPE cannot
func (parser *Parser) isStatementStart(tokenType token.TokenType) bool {
	switch tokenType {
	case token.LEKHA, token.DIA, token.DEIDE, token.JADI, token.GHURA, token.JETEBELEJAIN,
		token.KARYA, token.CHESTA, token.BAHARIPADE, token.CHADIDE,
		token.CONST, token.NUA, token.IDENT,
		token.SANKHYA, token.SABDA, token.AKSHARA, token.DASMIC, token.SATYA,
		token.KRAMA, token.MANA, token.THAKA, token.DHADHI:
		return true
	default:
		return false
	}
}

// syncAfterTopLevelStatement moves past a line end at the top level of a file.
// Example: after karya f() { } the next token is KARYA and must not be skipped
func (parser *Parser) syncAfterTopLevelStatement() {
	if parser.peekToken.Type == token.PIPE || parser.currentToken.Type == token.PIPE {
		parser.consumeStatementEnd()
		return
	}
	if parser.isStatementStart(parser.currentToken.Type) {
		return
	}
	parser.syncAfterBlockStatement()
}

// syncAfterBlockStatement moves past the | at the end of a statement inside a block.
// Example: after chadide | it moves forward to the next line in the block
func (parser *Parser) syncAfterBlockStatement() {
	tokenBeforeSync := parser.currentToken
	parser.consumeStatementEnd()
	if parser.currentToken != tokenBeforeSync {
		return
	}
	if parser.currentToken.Type == token.EOF || parser.currentToken.Type == token.RBRACE {
		return
	}
	parser.nextToken()
}

// consumeStatementEnd moves past the | at the end of a line when it is present.
// Example: after lekha("hi") the parser stops on ")" and this moves past "|"
func (parser *Parser) consumeStatementEnd() {
	if parser.peekToken.Type == token.PIPE {
		parser.nextToken()
	}
	if parser.currentToken.Type == token.PIPE {
		parser.nextToken()
	}
}

// parseBlock reads statements inside { ... } until the closing brace.
// Example: { lekha(x) | lekha(y) | } returns two LekhaStatement nodes
func (parser *Parser) parseBlock() []ast.Statement {
	var statements []ast.Statement
	for parser.currentToken.Type != token.RBRACE && parser.currentToken.Type != token.EOF {
		tokenBefore := parser.currentToken
		statement := parser.parseStatement()
		if statement != nil {
			statements = append(statements, statement)
		}
		parser.syncAfterBlockStatement()
		if parser.currentToken == tokenBefore &&
			parser.currentToken.Type != token.EOF &&
			parser.currentToken.Type != token.RBRACE {
			parser.nextToken()
		}
	}
	return statements
}

// parseExpressionList reads comma-separated expressions inside parentheses.
// Example: (a, b, c) returns three parsed expressions
func (parser *Parser) parseExpressionList(end token.TokenType) []ast.Expression {
	var expressions []ast.Expression
	if parser.peekToken.Type == end {
		parser.nextToken()
		return expressions
	}
	parser.nextToken()
	expressions = append(expressions, parser.parseExpression(lowestPrecedence))
	for parser.peekToken.Type == token.COMMA {
		parser.nextToken()
		parser.nextToken()
		expressions = append(expressions, parser.parseExpression(lowestPrecedence))
	}
	if !parser.expectPeek(end) {
		return expressions
	}
	return expressions
}

// parseIdentList reads comma-separated identifier names.
// Example: s, p, q returns ["s", "p", "q"]
func (parser *Parser) parseIdentList() []string {
	var names []string
	names = append(names, parser.currentToken.Literal)
	for parser.peekToken.Type == token.COMMA {
		parser.nextToken()
		parser.nextToken()
		names = append(names, parser.currentToken.Literal)
	}
	return names
}
