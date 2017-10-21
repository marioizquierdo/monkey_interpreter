package lexer_test

import (
	"testing"

	"github.com/marioizquierdo/monkey_interpreter/lib/lexer"
)

func Test_NextToken_FewChars(t *testing.T) {
	input := `=+(`
	l := lexer.New(input)

	tok1 := l.NextToken()
	exp1 := lexer.NewToken(lexer.ASSIGN, '=')
	if !lexer.IsTokenEq(tok1, exp1) {
		t.Fatalf("expected tok1: %+v to be: %+v", tok1, exp1)
	}

	tok2 := l.NextToken()
	exp2 := lexer.NewToken(lexer.PLUS, '+')
	if !lexer.IsTokenEq(tok2, exp2) {
		t.Fatalf("expected tok2: %+v to be: %+v", tok2, exp2)
	}

	tok3 := l.NextToken()
	exp3 := lexer.NewToken(lexer.LPAREN, '(')
	if !lexer.IsTokenEq(tok3, exp3) {
		t.Fatalf("expected tok3: %+v to be: %+v", tok3, exp3)
	}
}

func Test_NextToken_MoreChars(t *testing.T) {
	input := `=+(){},;`
	expectedTokens := []lexer.Token{
		{lexer.ASSIGN, "="},
		{lexer.PLUS, "+"},
		{lexer.LPAREN, "("},
		{lexer.RPAREN, ")"},
		{lexer.LBRACE, "{"},
		{lexer.RBRACE, "}"},
		{lexer.COMMA, ","},
		{lexer.SEMICOLON, ";"},
		{lexer.EOF, ""},
	}
	l := lexer.New(input)

	for i, expectedToken := range expectedTokens {
		tok := l.NextToken()
		if !lexer.IsTokenEq(tok, expectedToken) {
			t.Fatalf("expected tok[%d]: %+v to be: %+v", i, tok, expectedToken)
		}
	}
}
