package ast

// VarStatement declares a new variable with a type and optional start value.
// Example: sankhya x = 10 | becomes VarStatement{Name: "x", TypeName: "sankhya", Value: 10}
type VarStatement struct {
	Name     string
	TypeName string
	Value    Expression
	IsConst  bool
}

func (node *VarStatement) statementNode()       {}
func (node *VarStatement) TokenLiteral() string { return node.TypeName }

// AssignStatement gives new values to one or more existing variables.
// Example: s, p, q = misana(p, q) | sets three variables at once
type AssignStatement struct {
	Targets  []string
	LeftHand Expression
	Value    Expression
}

func (node *AssignStatement) statementNode()       {}
func (node *AssignStatement) TokenLiteral() string { return "=" }

// LekhaStatement prints values to the screen.
// Example: lekha("hi") | becomes LekhaStatement with one string argument
type LekhaStatement struct {
	Arguments []Expression
}

func (node *LekhaStatement) statementNode()       {}
func (node *LekhaStatement) TokenLiteral() string { return "lekha" }

// DiaStatement reads input from the user into variables.
// Example: dia(p >> q) | reads two numbers into p and q
type DiaStatement struct {
	Targets []string
}

func (node *DiaStatement) statementNode()       {}
func (node *DiaStatement) TokenLiteral() string { return "dia" }

// DeideStatement returns values from a function.
// Example: deide (a, b) | sends a and b back to the caller
type DeideStatement struct {
	Values []Expression
}

func (node *DeideStatement) statementNode()       {}
func (node *DeideStatement) TokenLiteral() string { return "deide" }

// IfStatement runs code only when a condition is true.
// Example: jadi x > 10 { lekha("big") | } has Condition x > 10
type IfStatement struct {
	Condition   Expression
	Consequence []Statement
	ElseIfs     []ElseIfClause
	Alternative []Statement
}

func (node *IfStatement) statementNode()       {}
func (node *IfStatement) TokenLiteral() string { return "jadi" }

// GhuraStatement is a for loop that counts from start to end.
// Example: ghura sankhya i = 0 -> 5 | i++ { ... } runs the body six times
type GhuraStatement struct {
	VarName  string
	TypeName string
	Start    Expression
	End      Expression
	Step     string
	Body     []Statement
}

func (node *GhuraStatement) statementNode()       {}
func (node *GhuraStatement) TokenLiteral() string { return "ghura" }

// JetebeleJainStatement is a while loop that runs while a condition is true.
// Example: jetebeleJain x > 0 { ... } keeps going until x is not greater than 0
type JetebeleJainStatement struct {
	Condition Expression
	Body      []Statement
}

func (node *JetebeleJainStatement) statementNode()       {}
func (node *JetebeleJainStatement) TokenLiteral() string { return "jetebeleJain" }

// BaharipadeStatement stops the nearest loop early.
// Example: baharipade | inside a ghura loop exits the loop
type BaharipadeStatement struct{}

func (node *BaharipadeStatement) statementNode()       {}
func (node *BaharipadeStatement) TokenLiteral() string { return "baharipade" }

// ChadideStatement skips to the next step of the nearest loop.
// Example: chadide | inside ghura jumps to the next i value
type ChadideStatement struct{}

func (node *ChadideStatement) statementNode()       {}
func (node *ChadideStatement) TokenLiteral() string { return "chadide" }

// KaryaStatement defines a named function with parameters and a body.
// Example: karya misana(pratham sankhya) (sankhya) { ... } is one function
type KaryaStatement struct {
	Name        string
	Parameters  []Parameter
	ReturnTypes []string
	Body        []Statement
}

func (node *KaryaStatement) statementNode()       {}
func (node *KaryaStatement) TokenLiteral() string { return "karya" }

// SreniStatement defines a class with fields and methods.
// Example: sreni Point { sankhya x ; karya sum() (sankhya) { ... } }
type SreniStatement struct {
	Name    string
	Fields  []*VarStatement
	Methods []*KaryaStatement
}

func (node *SreniStatement) statementNode()       {}
func (node *SreniStatement) TokenLiteral() string { return "sreni" }

// ChestaStatement runs code and catches errors in a dhare block.
// Example: chesta { ... } dhare { ... } tries the first block, handles errors in the second
type ChestaStatement struct {
	TryBody   []Statement
	CatchBody []Statement
}

func (node *ChestaStatement) statementNode()       {}
func (node *ChestaStatement) TokenLiteral() string { return "chesta" }

// AnaaStatement imports another module by name.
// Example: anaa fmt | brings in the fmt module
type AnaaStatement struct {
	Path string
}

func (node *AnaaStatement) statementNode()       {}
func (node *AnaaStatement) TokenLiteral() string { return "anaa" }

// ExpressionStatement wraps a bare expression used as a line of code.
// Example: misana(p, q) | alone on a line becomes ExpressionStatement
type ExpressionStatement struct {
	Expression Expression
}

func (node *ExpressionStatement) statementNode()       {}
func (node *ExpressionStatement) TokenLiteral() string { return node.Expression.TokenLiteral() }
