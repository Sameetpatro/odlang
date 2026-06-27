package parser

import (
	"strconv"

	"github.com/Sameetpatro/odlang/ast"
	"github.com/Sameetpatro/odlang/token"
)

const (
	lowestPrecedence = iota
	orPrecedence
	andPrecedence
	equalsPrecedence
	lessGreaterPrecedence
	sumPrecedence
	productPrecedence
	powerPrecedence
	prefixPrecedence
	callPrecedence
)

// precedences maps token types to binding strength for expression parsing.
// Example: * binds tighter than + so 1 + 2 * 3 groups as 1 + (2 * 3)
var precedences = map[token.TokenType]int{
	token.AAU:      orPrecedence,
	token.SAHITA:   andPrecedence,
	token.EQ:       equalsPrecedence,
	token.NOT_EQ:   equalsPrecedence,
	token.LT:       lessGreaterPrecedence,
	token.GT:       lessGreaterPrecedence,
	token.LTE:      lessGreaterPrecedence,
	token.GTE:      lessGreaterPrecedence,
	token.PLUS:     sumPrecedence,
	token.MINUS:    sumPrecedence,
	token.STAR:     productPrecedence,
	token.SLASH:    productPrecedence,
	token.EXP:      powerPrecedence,
	token.LPAREN:   callPrecedence,
	token.LBRACK:   callPrecedence,
	token.DOT:      callPrecedence,
}

// parseExpression reads one expression with operators at or above the given precedence.
// Example: parseExpression(lowest) reads the full 1 + 2 * 3 tree
func (parser *Parser) parseExpression(precedence int) ast.Expression {
	left := parser.parsePrefixExpression()
	for precedence < parser.peekPrecedence() {
		parser.nextToken()
		left = parser.parseInfixExpression(left)
	}
	return left
}

// peekPrecedence returns how tightly the next token binds.
// Example: peekPrecedence for + returns sumPrecedence
func (parser *Parser) peekPrecedence() int {
	if precedence, ok := precedences[parser.peekToken.Type]; ok {
		return precedence
	}
	return lowestPrecedence
}

// currentPrecedence returns binding strength for the current token.
// Example: currentPrecedence for * returns productPrecedence
func (parser *Parser) currentPrecedence() int {
	if precedence, ok := precedences[parser.currentToken.Type]; ok {
		return precedence
	}
	return lowestPrecedence
}

// parsePrefixExpression reads a value or operator-before-value form.
// Example: !han and -5 start with prefix operators
func (parser *Parser) parsePrefixExpression() ast.Expression {
	switch parser.currentToken.Type {
	case token.BANG:
		return parser.parsePrefixOperator("!")
	case token.MINUS:
		return parser.parsePrefixOperator("-")
	case token.INT:
		return parser.parseIntegerLiteral()
	case token.FLOAT:
		return parser.parseFloatLiteral()
	case token.STRING:
		return &ast.StringLiteral{Value: parser.currentToken.Literal}
	case token.CHAR:
		charValue := rune(0)
		if len(parser.currentToken.Literal) > 0 {
			charValue = rune(parser.currentToken.Literal[0])
		}
		return &ast.CharLiteral{Value: charValue}
	case token.HAN:
		return &ast.BooleanLiteral{Value: true}
	case token.NA:
		return &ast.BooleanLiteral{Value: false}
	case token.KHALI:
		return &ast.NullLiteral{}
	case token.IDENT:
		return parser.parseIdentifierExpression()
	case token.LPAREN:
		return parser.parseGroupedExpression()
	default:
		if isTypeToken(parser.currentToken.Type) && parser.peekToken.Type == token.LPAREN {
			return parser.parseTypeCastExpression()
		}
		return nil
	}
}

// parsePrefixOperator builds a PrefixExpression and parses the right side.
// Example: !other becomes PrefixExpression{Operator: "!", Right: other}
func (parser *Parser) parsePrefixOperator(operator string) ast.Expression {
	expression := &ast.PrefixExpression{Operator: operator}
	parser.nextToken()
	expression.Right = parser.parseExpression(prefixPrecedence)
	return expression
}

// parseGroupedExpression reads ( inner expression ).
// Example: (1 + 2) returns the InfixExpression for 1 + 2
func (parser *Parser) parseGroupedExpression() ast.Expression {
	parser.nextToken()
	expression := parser.parseExpression(lowestPrecedence)
	if !parser.expectPeek(token.RPAREN) {
		return expression
	}
	return expression
}

// parseIntegerLiteral converts the current INT token to an IntegerLiteral node.
// Example: token "10" becomes IntegerLiteral{Value: 10}
func (parser *Parser) parseIntegerLiteral() ast.Expression {
	value, err := strconv.ParseInt(parser.currentToken.Literal, 10, 64)
	if err != nil {
		parser.errors = append(parser.errors, "could not parse integer: "+parser.currentToken.Literal)
		return &ast.IntegerLiteral{}
	}
	return &ast.IntegerLiteral{Value: value}
}

// parseFloatLiteral converts the current FLOAT token to a FloatLiteral node.
// Example: token "3.14" becomes FloatLiteral{Value: 3.14}
func (parser *Parser) parseFloatLiteral() ast.Expression {
	value, err := strconv.ParseFloat(parser.currentToken.Literal, 64)
	if err != nil {
		parser.errors = append(parser.errors, "could not parse float: "+parser.currentToken.Literal)
		return &ast.FloatLiteral{}
	}
	return &ast.FloatLiteral{Value: value}
}

// parseIdentifierExpression reads a name and optional call or index suffix.
// Example: misana(p, q) becomes CallExpression after reading misana
func (parser *Parser) parseIdentifierExpression() ast.Expression {
	name := parser.currentToken.Literal
	if parser.peekToken.Type == token.LPAREN {
		return parser.parseCallExpression(name)
	}
	if parser.peekToken.Type == token.LBRACK {
		parser.nextToken()
		return parser.parseIndexExpression(&ast.Identifier{Name: name})
	}
	return &ast.Identifier{Name: name}
}

// parseTypeCastExpression reads sabda(x) style type conversion.
// Example: sabda(s) becomes TypeCastExpression{TargetType: "sabda", Value: s}
func (parser *Parser) parseTypeCastExpression() ast.Expression {
	targetType := parser.currentToken.Literal
	parser.nextToken()
	if parser.currentToken.Type != token.LPAREN {
		if !parser.expectPeek(token.LPAREN) {
			return &ast.TypeCastExpression{TargetType: targetType}
		}
	}
	parser.nextToken()
	value := parser.parseExpression(lowestPrecedence)
	if !parser.expectPeek(token.RPAREN) {
		return &ast.TypeCastExpression{TargetType: targetType, Value: value}
	}
	return &ast.TypeCastExpression{TargetType: targetType, Value: value}
}

// parseInfixExpression builds an infix node or handles call/index suffixes.
// Example: left + right becomes InfixExpression with Operator "+"
func (parser *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	switch parser.currentToken.Type {
	case token.LPAREN:
		if member, ok := left.(*ast.MemberExpression); ok {
			call := &ast.CallExpression{Receiver: member.Object, Function: member.Member}
			call.Arguments = parser.parseExpressionList(token.RPAREN)
			return call
		}
		if ident, ok := left.(*ast.Identifier); ok {
			expression := &ast.CallExpression{Function: ident.Name}
			expression.Arguments = parser.parseExpressionList(token.RPAREN)
			return expression
		}
	case token.DOT:
		member := &ast.MemberExpression{Object: left}
		parser.nextToken()
		member.Member = parser.currentToken.Literal
		return member
	case token.LBRACK:
		return parser.parseIndexExpression(left)
	case token.EXP:
		expression := &ast.InfixExpression{Left: left, Operator: parser.currentToken.Literal}
		parser.nextToken()
		expression.Right = parser.parseExpression(powerPrecedence)
		return expression
	}
	expression := &ast.InfixExpression{
		Left:     left,
		Operator: parser.currentToken.Literal,
	}
	precedence := parser.currentPrecedence()
	parser.nextToken()
	expression.Right = parser.parseExpression(precedence)
	return expression
}

// parseCallExpression reads functionName(arg1, arg2, ...) after the name.
// Example: misana(p, q) returns CallExpression with two arguments
func (parser *Parser) parseCallExpression(functionName string) ast.Expression {
	expression := &ast.CallExpression{Function: functionName}
	parser.nextToken()
	expression.Arguments = parser.parseExpressionList(token.RPAREN)
	return expression
}

// parseIndexExpression reads left[index] and may chain another index or call.
// Example: nums[0] becomes IndexExpression with Left nums and Index 0
func (parser *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
	expression := &ast.IndexExpression{Left: left}
	parser.nextToken()
	expression.Index = parser.parseExpression(lowestPrecedence)
	if !parser.expectPeek(token.RBRACK) {
		return expression
	}
	if parser.peekToken.Type == token.LBRACK {
		parser.nextToken()
		return parser.parseIndexExpression(expression)
	}
	if parser.peekToken.Type == token.LPAREN {
		if ident, ok := expression.Left.(*ast.Identifier); ok {
			call := &ast.CallExpression{Function: ident.Name}
			parser.nextToken()
			call.Arguments = parser.parseExpressionList(token.RPAREN)
			return call
		}
	}
	return expression
}
