package lexer_test

import (
	"testing"

	"github.com/marioizquierdo/monkey_interpreter/lib/lexer"
)

func Test_NextToken_Simple(t *testing.T) {
	input := `=+`
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
			t.Fatalf("expected tok[%d]: %v to be: %v", i, tok, expectedToken)
		}
	}
}

func Test_NextToken_LittleProgram(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y;
};

let result = add(five, ten);
!-/*5;
5 < 10 > 5;
`
	expectedTokens := []lexer.Token{
		{lexer.LET, "let"},
		{lexer.IDENT, "five"},
		{lexer.ASSIGN, "="},
		{lexer.INT, "5"},
		{lexer.SEMICOLON, ";"},
		{lexer.LET, "let"},
		{lexer.IDENT, "ten"},
		{lexer.ASSIGN, "="},
		{lexer.INT, "10"},
		{lexer.SEMICOLON, ";"},
		{lexer.LET, "let"},
		{lexer.IDENT, "add"},
		{lexer.ASSIGN, "="},
		{lexer.FUNCTION, "fn"},
		{lexer.LPAREN, "("},
		{lexer.IDENT, "x"},
		{lexer.COMMA, ","},
		{lexer.IDENT, "y"},
		{lexer.RPAREN, ")"},
		{lexer.LBRACE, "{"},
		{lexer.IDENT, "x"},
		{lexer.PLUS, "+"},
		{lexer.IDENT, "y"},
		{lexer.SEMICOLON, ";"},
		{lexer.RBRACE, "}"},
		{lexer.SEMICOLON, ";"},
		{lexer.LET, "let"},
		{lexer.IDENT, "result"},
		{lexer.ASSIGN, "="},
		{lexer.IDENT, "add"},
		{lexer.LPAREN, "("},
		{lexer.IDENT, "five"},
		{lexer.COMMA, ","},
		{lexer.IDENT, "ten"},
		{lexer.RPAREN, ")"},
		{lexer.SEMICOLON, ";"},
		{lexer.BANG, "!"},
		{lexer.MINUS, "-"},
		{lexer.SLASH, "/"},
		{lexer.ASTERISK, "*"},
		{lexer.INT, "5"},
		{lexer.SEMICOLON, ";"},
		{lexer.INT, "5"},
		{lexer.LT, "<"},
		{lexer.INT, "10"},
		{lexer.GT, ">"},
		{lexer.INT, "5"},
		{lexer.SEMICOLON, ";"},
		{lexer.EOF, ""},
	}
	l := lexer.New(input)

	for i, expectedToken := range expectedTokens {
		tok := l.NextToken()
		if !lexer.IsTokenEq(tok, expectedToken) {
			t.Fatalf("expected tok[%d]: %v to be: %v", i, tok, expectedToken)
		}
	}
}
