package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Sameetpatro/odlang/ast"
	"github.com/Sameetpatro/odlang/interpreter"
	"github.com/Sameetpatro/odlang/lexer"
	"github.com/Sameetpatro/odlang/parser"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	switch os.Args[1] {
	case "run":
		runFile()
	case "repl":
		runREPL()
	case "version":
		fmt.Println("OdLang v1.0.0")
	case "help":
		printHelp()
	default:
		printHelp()
	}
}

func runFile() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: od run <file.od>")
		os.Exit(1)
	}

	source, err := os.ReadFile(os.Args[2])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	program, ok := parseSource(string(source))
	if !ok {
		os.Exit(1)
	}

	interpreter.Eval(program)
}

func runREPL() {
	fmt.Println("OdLang REPL v1.0 | type 'bahara' to exit")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">>> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		if line == "bahara" {
			break
		}
		if line == "" {
			continue
		}

		program, ok := parseSource(line)
		if !ok {
			continue
		}

		result := interpreter.EvalProgram(program)
		if result != nil {
			fmt.Println(formatREPLResult(result))
		}
	}
}

func parseSource(source string) (*ast.Program, bool) {
	lexerInstance := lexer.New(source)
	parserInstance := parser.New(lexerInstance)
	program := parserInstance.ParseProgram()

	if len(parserInstance.Errors()) > 0 {
		for _, parseError := range parserInstance.Errors() {
			fmt.Println("parse error:", parseError)
		}
		return nil, false
	}
	return program, true
}

func formatREPLResult(value interface{}) string {
	if value == nil {
		return "khali"
	}
	switch v := value.(type) {
	case int64:
		return fmt.Sprintf("%d", v)
	case float64:
		return fmt.Sprintf("%g", v)
	case string:
		return v
	case bool:
		if v {
			return "han"
		}
		return "na"
	default:
		return fmt.Sprintf("%v", v)
	}
}

func printHelp() {
	fmt.Println("OdLang — tree-walk interpreter")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  od run <file.od>   Run an OdLang program")
	fmt.Println("  od repl            Start interactive REPL")
	fmt.Println("  od version         Show version")
	fmt.Println("  od help            Show this help")
}
