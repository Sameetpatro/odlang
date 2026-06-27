package parser

import (
	"github.com/Sameetpatro/odlang/ast"
	"github.com/Sameetpatro/odlang/token"
)

// parseVarStatement reads a typed variable declaration like sankhya x = 10 |
// Example: krama nums(5, 0) | stores a call expression as the initial value
func (parser *Parser) parseVarStatement(typeName string, isConst bool) *ast.VarStatement {
	statement := &ast.VarStatement{IsConst: isConst}
	if parser.currentToken.Type == token.NUA {
		statement.TypeName = "nua"
		parser.nextToken()
	} else {
		statement.TypeName = typeName
		if statement.TypeName == "" {
			statement.TypeName = parser.currentToken.Literal
		}
		parser.nextToken()
	}
	if !parser.expectPeek(token.IDENT) {
		return statement
	}
	statement.Name = parser.currentToken.Literal
	if parser.peekToken.Type == token.ASSIGN {
		parser.nextToken()
		parser.nextToken()
		statement.Value = parser.parseExpression(lowestPrecedence)
	} else if parser.peekToken.Type == token.LPAREN {
		parser.nextToken()
		statement.Value = parser.parseCallExpression(statement.Name)
	}
	return statement
}

// parseConstStatement reads a constant declaration like const PI = 3.14 |
// Example: const MAX = 100 | sets IsConst true and stores the value
func (parser *Parser) parseConstStatement() *ast.VarStatement {
	parser.nextToken()
	statement := &ast.VarStatement{IsConst: true, Name: parser.currentToken.Literal}
	if parser.peekToken.Type == token.ASSIGN {
		parser.nextToken()
		parser.nextToken()
		statement.Value = parser.parseExpression(lowestPrecedence)
	}
	return statement
}

// parseIdentStatement reads assignment or a bare expression starting with a name.
// Example: nums[0] = 1 | becomes AssignStatement with an index on the left side
func (parser *Parser) parseIdentStatement() ast.Statement {
	name := parser.currentToken.Literal
	if parser.peekToken.Type == token.LBRACK {
		parser.nextToken()
		parser.nextToken()
		index := parser.parseExpression(lowestPrecedence)
		if !parser.expectPeek(token.RBRACK) {
			return nil
		}
		if parser.peekToken.Type == token.ASSIGN {
			parser.nextToken()
			parser.nextToken()
			return &ast.AssignStatement{
				LeftHand: &ast.IndexExpression{
					Left:  &ast.Identifier{Name: name},
					Index: index,
				},
				Value: parser.parseExpression(lowestPrecedence),
			}
		}
		return &ast.ExpressionStatement{
			Expression: &ast.IndexExpression{
				Left:  &ast.Identifier{Name: name},
				Index: index,
			},
		}
	}
	if parser.peekToken.Type == token.ASSIGN {
		return parser.parseAssignStatement([]string{name})
	}
	if parser.peekToken.Type == token.COMMA {
		targets := parser.parseIdentList()
		parser.nextToken()
		return parser.parseAssignStatement(targets)
	}
	return parser.parseExpressionStatement()
}

// parseAssignStatement reads x = expr or s, p, q = expr |
// Example: msg = lekhaSabda("hi") | assigns the call result to msg
func (parser *Parser) parseAssignStatement(targets []string) *ast.AssignStatement {
	parser.nextToken()
	parser.nextToken()
	return &ast.AssignStatement{
		Targets: targets,
		Value:   parser.parseExpression(lowestPrecedence),
	}
}

// parseLekhaStatement reads lekha(arg1, arg2, ...) |
// Example: lekha("hello") | stores one string argument
func (parser *Parser) parseLekhaStatement() *ast.LekhaStatement {
	statement := &ast.LekhaStatement{}
	parser.nextToken()
	statement.Arguments = parser.parseExpressionList(token.RPAREN)
	return statement
}

// parseDiaStatement reads dia(p >> q) |
// Example: dia(p >> q) | stores Targets ["p", "q"]
func (parser *Parser) parseDiaStatement() *ast.DiaStatement {
	statement := &ast.DiaStatement{}
	parser.nextToken()
	if parser.currentToken.Type != token.LPAREN {
		if !parser.expectPeek(token.LPAREN) {
			return statement
		}
	}
	parser.nextToken()
	statement.Targets = append(statement.Targets, parser.currentToken.Literal)
	for parser.peekToken.Type == token.RSHIFT {
		parser.nextToken()
		parser.nextToken()
		statement.Targets = append(statement.Targets, parser.currentToken.Literal)
	}
	parser.expectPeek(token.RPAREN)
	return statement
}

// parseDeideStatement reads deide (a, b, c) |
// Example: deide (pratham + dutiya, pratham, dutiya) | stores three expressions
func (parser *Parser) parseDeideStatement() *ast.DeideStatement {
	parser.nextToken()
	return &ast.DeideStatement{Values: parser.parseExpressionList(token.RPAREN)}
}

// parseBaharipadeStatement reads baharipade |
// Example: baharipade | inside a loop becomes BaharipadeStatement
func (parser *Parser) parseBaharipadeStatement() *ast.BaharipadeStatement {
	parser.nextToken()
	return &ast.BaharipadeStatement{}
}

// parseChadideStatement reads chadide |
// Example: chadide | skips to the next loop step
func (parser *Parser) parseChadideStatement() *ast.ChadideStatement {
	parser.nextToken()
	return &ast.ChadideStatement{}
}

// parseExpressionStatement reads a bare expression used as a full line.
// Example: misana(p, q) | alone becomes ExpressionStatement wrapping the call
func (parser *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	return &ast.ExpressionStatement{Expression: parser.parseExpression(lowestPrecedence)}
}
