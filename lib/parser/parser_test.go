package parser

import (
	"fmt"
	"testing"

	"github.com/marioizquierdo/monkey_interpreter/lib/ast"
	"github.com/marioizquierdo/monkey_interpreter/lib/lexer"
)

func Test_Let_Statements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if errMsg := checkParserErrors(t, p); errMsg != "" {
		t.Fatal(errMsg)
	}

	if program == nil {
		t.Fatalf("expected ParseProgram() to be nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("expected len(ParseProgram()) to be 3, but is %d", len(program.Statements))
	}

	expectedIdentifiers := []string{"x", "y", "foobar"}
	for i, expectedIdentifier := range expectedIdentifiers {
		stmt := program.Statements[i]
		if errMsg := testLetStatement(t, stmt, expectedIdentifier); errMsg != "" {
			t.Fatalf(errMsg)
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) string {
	if s.TokenLiteral() != "let" {
		return fmt.Sprintf("expected s.TokenLiteral to be 'let', but is %q", s.TokenLiteral())
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		return fmt.Sprintf("expected s to be a *ast.LetStatement, but is a %T", s)
	}

	if letStmt.Name.Value != name {
		return fmt.Sprintf("expected letStmt.Name.Value to be %q, but is %q", name, letStmt.Name.Value)
	}

	if letStmt.Name.TokenLiteral() != name {
		return fmt.Sprintf("expected letStmt.Name to be %q, but is %q", name, letStmt.Name.TokenLiteral())
	}

	return ""
}

// checkParserErrors returns "" if no errors, but prints one error per line if it finds errors on the parser.
func checkParserErrors(t *testing.T, p *Parser) string {
	errMsgs := p.Errors()
	if len(errMsgs) != 0 {
		msg := fmt.Sprintf("parser has %d errors:\n", len(errMsgs))
		for _, errMsg := range errMsgs {
			msg += fmt.Sprintf("parser error: %s\n", errMsg)
		}
		return msg
	}

	return "" // ok
}
