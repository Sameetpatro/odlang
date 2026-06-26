package ast

// IntegerLiteral holds a whole number from the source code.
// Example: 10 in sankhya x = 10 is stored as IntegerLiteral{Value: 10}
type IntegerLiteral struct {
	Value int64
}

func (node *IntegerLiteral) expressionNode()      {}
func (node *IntegerLiteral) TokenLiteral() string { return "INT" }

// FloatLiteral holds a decimal number from the source code.
// Example: 3.14 in const PI = 3.14 is stored as FloatLiteral{Value: 3.14}
type FloatLiteral struct {
	Value float64
}

func (node *FloatLiteral) expressionNode()      {}
func (node *FloatLiteral) TokenLiteral() string { return "FLOAT" }

// StringLiteral holds text from inside double quotes.
// Example: "hello" in lekha("hello") is StringLiteral{Value: "hello"}
type StringLiteral struct {
	Value string
}

func (node *StringLiteral) expressionNode()      {}
func (node *StringLiteral) TokenLiteral() string { return "STRING" }

// CharLiteral holds one character from inside single quotes.
// Example: 'a' in akshara c = 'a' is CharLiteral{Value: 'a'}
type CharLiteral struct {
	Value rune
}

func (node *CharLiteral) expressionNode()      {}
func (node *CharLiteral) TokenLiteral() string { return "CHAR" }

// BooleanLiteral holds true or false from han or na.
// Example: han in satya ok = han is BooleanLiteral{Value: true}
type BooleanLiteral struct {
	Value bool
}

func (node *BooleanLiteral) expressionNode()      {}
func (node *BooleanLiteral) TokenLiteral() string { return "BOOL" }

// NullLiteral represents the khali (null) value.
// Example: nua x = khali uses NullLiteral on the right side
type NullLiteral struct{}

func (node *NullLiteral) expressionNode()      {}
func (node *NullLiteral) TokenLiteral() string { return "khali" }

// Identifier is a variable or function name used in an expression.
// Example: x in lekha(x) is Identifier{Name: "x"}
type Identifier struct {
	Name string
}

func (node *Identifier) expressionNode()      {}
func (node *Identifier) TokenLiteral() string { return node.Name }

// InfixExpression is a math or compare operation with left and right sides.
// Example: pratham + dutiya is InfixExpression with Operator "+"
type InfixExpression struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (node *InfixExpression) expressionNode()      {}
func (node *InfixExpression) TokenLiteral() string { return node.Operator }

// PrefixExpression is an operator placed before one value.
// Example: !other is PrefixExpression with Operator "!"
type PrefixExpression struct {
	Operator string
	Right    Expression
}

func (node *PrefixExpression) expressionNode()      {}
func (node *PrefixExpression) TokenLiteral() string { return node.Operator }

// CallExpression is a function call with a name and argument list.
// Example: misana(p, q) is CallExpression{Function: "misana", Arguments: [...]}
type CallExpression struct {
	Function  string
	Arguments []Expression
}

func (node *CallExpression) expressionNode()      {}
func (node *CallExpression) TokenLiteral() string { return node.Function }

// IndexExpression reads one item from an array by position.
// Example: nums[0] is IndexExpression with Left nums and Index 0
type IndexExpression struct {
	Left  Expression
	Index Expression
}

func (node *IndexExpression) expressionNode()      {}
func (node *IndexExpression) TokenLiteral() string { return "[" }

// TypeCastExpression converts a value to another type.
// Example: sabda(s) turns number s into a string
type TypeCastExpression struct {
	TargetType string
	Value      Expression
}

func (node *TypeCastExpression) expressionNode()      {}
func (node *TypeCastExpression) TokenLiteral() string { return node.TargetType }
