package parser_test

import (
	"os"
	"testing"
	"time"

	"github.com/Sameetpatro/odlang/ast"
	"github.com/Sameetpatro/odlang/lexer"
	"github.com/Sameetpatro/odlang/parser"
)

func TestMinimalDeide(t *testing.T) {
	source := "karya f() () { deide (1) ; }\n"
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
	source := "karya m() () {\n    deide (pratham + dutiya, pratham, dutiya) ;\n}\n"
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
    deide (pratham + dutiya, pratham, dutiya) ;
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

func TestGhuraChadideParse(t *testing.T) {
	cases := map[string]string{
		"simple ghura": `karya f() () {
    ghura sankhya i = 0 -> 5 ; i++ {
        lekha("loop") ;
    }
}`,
		"ghura with jadi/chadide": `karya f() () {
    ghura sankhya i = 0 -> 5 ; i++ {
        jadi i == 3 {
            chadide ;
        } ;
        lekha("loop") ;
    }
}`,
	}
	for name, source := range cases {
		t.Run(name, func(t *testing.T) {
			p := parser.New(lexer.New(source))
			program := p.ParseProgram()
			if len(p.Errors()) > 0 {
				t.Fatalf("errors: %v", p.Errors())
			}
			if len(program.Statements) != 1 {
				t.Fatalf("want 1 statement, got %d", len(program.Statements))
			}
		})
	}
}

func TestParseLexerTestTail(t *testing.T) {
	source := `karya aarambha() () {
    sankhya x = 10 ;
    jetebeleJain x > 0 {
        lekha("countdown: " + sabda(x)) ;
        x = x - 1 ;
        jadi x == 5 {
            baharipade ;
        } ;
    }

    dasmik ratio = PI / 2.0 ;
    akshara mark = '!' ;
    nua extra = khali ;
    lekha("ratio = " + sabda(ratio) + " mark = " + sabda(mark) + " extra = " + sabda(extra)) ;

    krama nums(5, 0) ;
    nums[0] = 1 ;
    nums[1] = 2 ;
    lekha("first cell: " + sabda(nums[0]) + " power: " + sabda(2**3)) ;

    sabda casted = sabda(s) ;
    lekha("casted sum: " + casted) ;

    chesta {
        lekha("try block ok") ;
        sankhya bad = 10 / 0 ;
    } dhare {
        lekha("catch block ran") ;
    }

    lekha("=== done ===") ;
}
`
	done := make(chan struct{}, 1)
	go func() {
		parser.New(lexer.New(source)).ParseProgram()
		done <- struct{}{}
	}()
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("parse hung")
	}
}

func TestParseLexerTestFile(t *testing.T) {
	source, err := os.ReadFile("../example/lexer_test.od")
	if err != nil {
		t.Fatalf("read lexer_test.od: %v", err)
	}
	done := make(chan []string, 1)
	go func() {
		p := parser.New(lexer.New(string(source)))
		p.ParseProgram()
		done <- p.Errors()
	}()
	select {
	case errs := <-done:
		if len(errs) > 0 {
			t.Fatalf("errors: %v", errs)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("parse hung")
	}
}
