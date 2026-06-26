package ast

// Node is the base type for every piece of the syntax tree.
// Example: both "x = 10" and "lekha(x)" are nodes in the tree
type Node interface {
	TokenLiteral() string
}

// Statement is a line of code that does an action but has no value.
// Example: lekha("hi") | is a statement, not an expression
type Statement interface {
	Node
	statementNode()
}

// Expression is code that produces a value.
// Example: 10 + 20 is an expression with value 30
type Expression interface {
	Node
	expressionNode()
}

// Program is the root of the whole tree and holds all top-level statements.
// Example: a file with karya aarambha() { ... } has one Program node
type Program struct {
	Statements []Statement
}

// TokenLiteral returns the first statement keyword for quick checks.
// Example: a program starting with karya returns "karya"
func (program *Program) TokenLiteral() string {
	if len(program.Statements) > 0 {
		return program.Statements[0].TokenLiteral()
	}
	return ""
}

// Parameter is one function input with a name and type.
// Example: pratham sankhya in karya misana(pratham sankhya) is one Parameter
type Parameter struct {
	Name     string
	TypeName string
}

// ElseIfClause is one extra branch in an if statement.
// Example: nahele jadi x == 10 { ... } adds one ElseIfClause
type ElseIfClause struct {
	Condition Expression
	Body      []Statement
}
