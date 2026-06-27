package ast

import (
	"fmt"
	"strings"
)

// String turns the whole program into readable text for debugging.
// Example: printing a Program shows each statement on its own line
func (program *Program) String() string {
	var parts []string
	for _, statement := range program.Statements {
		parts = append(parts, statementString(statement))
	}
	return strings.Join(parts, "\n")
}

// statementString formats one statement node as text.
// Example: a LekhaStatement becomes lekha(...)
func statementString(statement Statement) string {
	switch node := statement.(type) {
	case *VarStatement:
		valueText := "nil"
		if node.Value != nil {
			valueText = expressionString(node.Value)
		}
		constText := ""
		if node.IsConst {
			constText = "const "
		}
		return fmt.Sprintf("%s%s %s = %s", constText, node.TypeName, node.Name, valueText)
	case *AssignStatement:
		if node.LeftHand != nil {
			return fmt.Sprintf("%s = %s", expressionString(node.LeftHand), expressionString(node.Value))
		}
		return fmt.Sprintf("%s = %s", strings.Join(node.Targets, ", "), expressionString(node.Value))
	case *LekhaStatement:
		return fmt.Sprintf("lekha(%s)", joinExpressions(node.Arguments))
	case *DiaStatement:
		return fmt.Sprintf("dia(%s)", strings.Join(node.Targets, " >> "))
	case *DeideStatement:
		return fmt.Sprintf("deide (%s)", joinExpressions(node.Values))
	case *IfStatement:
		return fmt.Sprintf("jadi %s { ... }", expressionString(node.Condition))
	case *GhuraStatement:
		return fmt.Sprintf("ghura %s %s = %s -> %s | %s { ... }",
			node.TypeName, node.VarName, expressionString(node.Start),
			expressionString(node.End), node.Step)
	case *JetebeleJainStatement:
		return fmt.Sprintf("jetebeleJain %s { ... }", expressionString(node.Condition))
	case *BaharipadeStatement:
		return "baharipade"
	case *ChadideStatement:
		return "chadide"
	case *KaryaStatement:
		return karyaString(node)
	case *SreniStatement:
		return fmt.Sprintf("sreni %s { %d fields, %d methods }", node.Name, len(node.Fields), len(node.Methods))
	case *ChestaStatement:
		return "chesta { ... } dhare { ... }"
	case *AnaaStatement:
		return fmt.Sprintf("anaa %s", node.Path)
	case *ExpressionStatement:
		return expressionString(node.Expression)
	default:
		return fmt.Sprintf("%T", statement)
	}
}

// karyaString formats a function definition with params and return types.
// Example: karya misana(pratham sankhya) (sankhya) { ... }
func karyaString(node *KaryaStatement) string {
	paramParts := make([]string, len(node.Parameters))
	for index, param := range node.Parameters {
		paramParts[index] = fmt.Sprintf("%s %s", param.Name, param.TypeName)
	}
	return fmt.Sprintf("karya %s (%s) (%s) { %d statements }",
		node.Name, strings.Join(paramParts, ", "),
		strings.Join(node.ReturnTypes, ", "), len(node.Body))
}

// expressionString formats one expression node as text.
// Example: pratham + dutiya becomes (pratham + dutiya)
func expressionString(expression Expression) string {
	switch node := expression.(type) {
	case *IntegerLiteral:
		return fmt.Sprintf("%d", node.Value)
	case *FloatLiteral:
		return fmt.Sprintf("%g", node.Value)
	case *StringLiteral:
		return fmt.Sprintf("%q", node.Value)
	case *CharLiteral:
		return fmt.Sprintf("'%c'", node.Value)
	case *BooleanLiteral:
		return fmt.Sprintf("%t", node.Value)
	case *NullLiteral:
		return "khali"
	case *Identifier:
		return node.Name
	case *InfixExpression:
		return fmt.Sprintf("(%s %s %s)", expressionString(node.Left), node.Operator, expressionString(node.Right))
	case *PrefixExpression:
		return fmt.Sprintf("(%s%s)", node.Operator, expressionString(node.Right))
	case *CallExpression:
		if node.Receiver != nil {
			return fmt.Sprintf("%s.%s(%s)", expressionString(node.Receiver), node.Function, joinExpressions(node.Arguments))
		}
		return fmt.Sprintf("%s(%s)", node.Function, joinExpressions(node.Arguments))
	case *MemberExpression:
		return fmt.Sprintf("%s.%s", expressionString(node.Object), node.Member)
	case *IndexExpression:
		return fmt.Sprintf("%s[%s]", expressionString(node.Left), expressionString(node.Index))
	case *TypeCastExpression:
		return fmt.Sprintf("%s(%s)", node.TargetType, expressionString(node.Value))
	default:
		return fmt.Sprintf("%T", expression)
	}
}

// joinExpressions joins expression strings with commas.
// Example: [a, b, c] becomes "a, b, c"
func joinExpressions(expressions []Expression) string {
	parts := make([]string, len(expressions))
	for index, expression := range expressions {
		parts[index] = expressionString(expression)
	}
	return strings.Join(parts, ", ")
}
