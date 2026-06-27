package parser

import (
	"github.com/Sameetpatro/odlang/ast"
	"github.com/Sameetpatro/odlang/token"
)

// parseIfStatement reads jadi / nahele jadi / nahele blocks.
// Example: jadi s > 100 { ... } nahele { ... } builds IfStatement with alternative
func (parser *Parser) parseIfStatement() *ast.IfStatement {
	statement := &ast.IfStatement{}
	parser.nextToken()
	statement.Condition = parser.parseExpression(lowestPrecedence)
	statement.Consequence = parser.parseBlockBody()
	parser.parseElseParts(statement)
	return statement
}

// parseBlockBody reads { statements } and consumes the closing brace.
// Example: { lekha(x) | } returns statements and leaves parser past RBRACE
func (parser *Parser) parseBlockBody() []ast.Statement {
	if parser.currentToken.Type != token.LBRACE {
		if !parser.expectPeek(token.LBRACE) {
			return nil
		}
	}
	parser.nextToken()
	statements := parser.parseBlock()
	if parser.currentToken.Type == token.RBRACE {
		parser.nextToken()
	}
	return statements
}

// parseElseParts adds nahele jadi and nahele branches to an if statement.
// Example: nahele jadi x == 10 { ... } adds one ElseIfClause
func (parser *Parser) parseElseParts(statement *ast.IfStatement) {
	for parser.currentToken.Type == token.NAHELE {
		parser.nextToken()
		if parser.currentToken.Type == token.JADI {
			clause := ast.ElseIfClause{}
			parser.nextToken()
			clause.Condition = parser.parseExpression(lowestPrecedence)
			clause.Body = parser.parseBlockBody()
			statement.ElseIfs = append(statement.ElseIfs, clause)
			continue
		}
		statement.Alternative = parser.parseBlockBody()
		break
	}
}

// parseGhuraStatement reads ghura sankhya i = 0 -> 5 ; i++ { ... }
// Example: ghura loop with i++ stores Step as "++" after the variable name
func (parser *Parser) parseGhuraStatement() *ast.GhuraStatement {
	statement := &ast.GhuraStatement{}
	parser.nextToken()
	statement.TypeName = parser.currentToken.Literal
	parser.nextToken()
	statement.VarName = parser.currentToken.Literal
	if !parser.expectPeek(token.ASSIGN) {
		return statement
	}
	parser.nextToken()
	statement.Start = parser.parseExpression(lowestPrecedence)
	if !parser.expectPeek(token.ARROW) {
		return statement
	}
	parser.nextToken()
	statement.End = parser.parseExpression(lowestPrecedence)
	if parser.peekToken.Type == token.SEMI {
		parser.nextToken()
		parser.nextToken()
	} else if parser.currentToken.Type == token.SEMI {
		parser.nextToken()
	}
	stepName := parser.currentToken.Literal
	statement.Step = stepName
	if parser.peekToken.Type == token.INCREMENT || parser.peekToken.Type == token.DECREMENT {
		statement.Step += parser.peekToken.Literal
		parser.nextToken()
	}
	statement.Body = parser.parseBlockBody()
	return statement
}

// parseJetebeleJainStatement reads jetebeleJain condition { body }
// Example: jetebeleJain x > 0 { lekha(x) | } stores condition and body
func (parser *Parser) parseJetebeleJainStatement() *ast.JetebeleJainStatement {
	statement := &ast.JetebeleJainStatement{}
	parser.nextToken()
	statement.Condition = parser.parseExpression(lowestPrecedence)
	statement.Body = parser.parseBlockBody()
	return statement
}

// parseKaryaStatement reads a full function definition with params and return types.
// Example: karya misana(pratham sankhya) (sankhya) { ... } builds KaryaStatement
func (parser *Parser) parseKaryaStatement() *ast.KaryaStatement {
	statement := &ast.KaryaStatement{}
	parser.nextToken()
	statement.Name = parser.currentToken.Literal
	statement.Parameters = parser.parseParameterList()
	if parser.peekToken.Type == token.LPAREN {
		statement.ReturnTypes = parser.parseReturnTypeList()
	}
	statement.Body = parser.parseBlockBody()
	return statement
}

// parseParameterList reads (name type, name type) after a function name.
// Example: (pratham sankhya, dutiya sankhya) returns two Parameter nodes
func (parser *Parser) parseParameterList() []ast.Parameter {
	var parameters []ast.Parameter
	if !parser.expectPeek(token.LPAREN) {
		return parameters
	}
	if parser.peekToken.Type == token.RPAREN {
		parser.nextToken()
		return parameters
	}
	parser.nextToken()
	parameters = append(parameters, parser.parseOneParameter())
	for parser.peekToken.Type == token.COMMA {
		parser.nextToken()
		parser.nextToken()
		parameters = append(parameters, parser.parseOneParameter())
	}
	parser.expectPeek(token.RPAREN)
	return parameters
}

// parseOneParameter reads one name and type pair like pratham sankhya.
// Example: dutiya sankhya becomes Parameter{Name: "dutiya", TypeName: "sankhya"}
func (parser *Parser) parseOneParameter() ast.Parameter {
	param := ast.Parameter{Name: parser.currentToken.Literal}
	parser.nextToken()
	param.TypeName = parser.currentToken.Literal
	return param
}

// parseReturnTypeList reads (sankhya, sankhya) return type names.
// Example: (sankhya, sankhya, sankhya) returns three type strings
func (parser *Parser) parseReturnTypeList() []string {
	var returnTypes []string
	if !parser.expectPeek(token.LPAREN) {
		return returnTypes
	}
	if parser.peekToken.Type == token.RPAREN {
		parser.nextToken()
		return returnTypes
	}
	parser.nextToken()
	returnTypes = append(returnTypes, parser.currentToken.Literal)
	for parser.peekToken.Type == token.COMMA {
		parser.nextToken()
		parser.nextToken()
		returnTypes = append(returnTypes, parser.currentToken.Literal)
	}
	parser.expectPeek(token.RPAREN)
	return returnTypes
}

// parseSreniStatement reads a class with field and method declarations.
// Example: sreni Point { sankhya x ; karya sum() (sankhya) { deide (x) ; } }
func (parser *Parser) parseSreniStatement() *ast.SreniStatement {
	statement := &ast.SreniStatement{}
	parser.nextToken()
	statement.Name = parser.currentToken.Literal
	if !parser.expectPeek(token.LBRACE) {
		return statement
	}
	parser.nextToken()
	for parser.currentToken.Type != token.RBRACE && parser.currentToken.Type != token.EOF {
		countBefore := parser.tokenCount
		switch parser.currentToken.Type {
		case token.KARYA:
			statement.Methods = append(statement.Methods, parser.parseKaryaStatement())
		case token.SANKHYA, token.SABDA, token.AKSHARA, token.DASMIC, token.SATYA,
			token.KRAMA, token.MANA, token.THAKA, token.DHADHI:
			statement.Fields = append(statement.Fields, parser.parseVarStatement(parser.currentToken.Literal, false))
		default:
			parser.errors = append(parser.errors,
				"expected field or karya inside sreni, got "+string(parser.currentToken.Type))
		}
		parser.syncAfterBlockStatement()
		if parser.tokenCount == countBefore &&
			parser.currentToken.Type != token.EOF &&
			parser.currentToken.Type != token.RBRACE &&
			!parser.isStatementStart(parser.currentToken.Type) {
			parser.nextToken()
		}
	}
	if parser.currentToken.Type == token.RBRACE {
		parser.nextToken()
	}
	return statement
}

// parseChestaStatement reads chesta { ... } dhare { ... }
// Example: try/catch blocks map to TryBody and CatchBody statement lists
func (parser *Parser) parseChestaStatement() *ast.ChestaStatement {
	statement := &ast.ChestaStatement{}
	parser.nextToken()
	statement.TryBody = parser.parseBlockBody()
	if parser.currentToken.Type == token.DHARE {
		parser.nextToken()
		statement.CatchBody = parser.parseBlockBody()
	}
	return statement
}
