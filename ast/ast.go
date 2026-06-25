package ast

// Node is implemented by every AST node.
// ast hauchi abstract syntax tree and aeita ra kama hauchi tree baneiba for each statements
type Node interface {
    TokenLiteral() string
}

// Statement nodes do not produce a value.
type Statement interface {
    Node
    statementNode()
}

// Program is the root node of every OdLang AST.
//aeithi sabu statement save hue
type Program struct {
    Statements []Statement
}

func (p *Program) TokenLiteral() string {
    if len(p.Statements) > 0 {
        return p.Statements[0].TokenLiteral()
    }
    return ""
}

// LekhaStatement represents: lekha("...")
type LekhaStatement struct {
    Value string // the string argument
}

func (ls *LekhaStatement) statementNode()       {}
func (ls *LekhaStatement) TokenLiteral() string { return "lekha" }