package parser_test

import (
	"testing"

	"github.com/Sameetpatro/odlang/ast"
	"github.com/Sameetpatro/odlang/lexer"
	"github.com/Sameetpatro/odlang/parser"
)

func TestMinimalDeide(t *testing.T) {
	source := "karya f() () { deide (1) | }\n"
	lexerInstance := lexer.New(source)
	parserInstance := parser.New(lexerInstance)
	program := parserInstance.ParseProgram()
	if len(parserInstance.Errors()) > 0 {
		t.Fatalf("errors: %v", parserInstance.Errors())
	}
	if len(program.Statements) != 1 {
		t.Fatalf("want 1 statement got %d", len(program.Statements))
	}
}

func TestDeideCommaList(t *testing.T) {
	source := "karya m() () {\n    deide (pratham + dutiya, pratham, dutiya) |\n}\n"
	lexerInstance := lexer.New(source)
	parserInstance := parser.New(lexerInstance)
	program := parserInstance.ParseProgram()
	if len(parserInstance.Errors()) > 0 {
		t.Fatalf("errors: %v", parserInstance.Errors())
	}
	if len(program.Statements) != 1 {
		t.Fatalf("want 1 statement got %d", len(program.Statements))
	}
}

func TestMisanaParse(t *testing.T) {
	source := `karya misana (pratham sankhya, dutiya sankhya) (sankhya, sankhya, sankhya) {
    deide (pratham + dutiya, pratham, dutiya) |
}`
	lexerInstance := lexer.New(source)
	parserInstance := parser.New(lexerInstance)
	program := parserInstance.ParseProgram()
	if len(parserInstance.Errors()) > 0 {
		t.Fatalf("errors: %v", parserInstance.Errors())
	}
	karyaStatement, ok := program.Statements[0].(*ast.KaryaStatement)
	if !ok {
		t.Fatalf("expected KaryaStatement, got %T", program.Statements[0])
	}
	if karyaStatement.Name != "misana" {
		t.Fatalf("expected name misana, got %q", karyaStatement.Name)
	}
	if len(karyaStatement.Parameters) != 2 {
		t.Fatalf("expected 2 parameters, got %d", len(karyaStatement.Parameters))
	}
	if karyaStatement.Parameters[0].Name != "pratham" || karyaStatement.Parameters[0].TypeName != "sankhya" {
		t.Fatalf("unexpected first parameter: %+v", karyaStatement.Parameters[0])
	}
	if len(karyaStatement.ReturnTypes) != 3 {
		t.Fatalf("expected 3 return types, got %d", len(karyaStatement.ReturnTypes))
	}
	deideStatement, ok := karyaStatement.Body[0].(*ast.DeideStatement)
	if !ok {
		t.Fatalf("expected DeideStatement body, got %T", karyaStatement.Body[0])
	}
	if len(deideStatement.Values) != 3 {
		t.Fatalf("expected 3 return values, got %d", len(deideStatement.Values))
	}
	infixExpression, ok := deideStatement.Values[0].(*ast.InfixExpression)
	if !ok {
		t.Fatalf("expected InfixExpression as first return value, got %T", deideStatement.Values[0])
	}
	if _, ok := infixExpression.Left.(*ast.Identifier); !ok {
		t.Fatalf("expected left identifier in infix expression")
	}
}
