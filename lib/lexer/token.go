package lexer

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Literals
	INT = "INT" // 1, 2, 333

	// Variable names
	IDENT = "IDENT"

	// Operators
	ASSIGN = "ASSIGN"
	PLUS   = "PLUS"

	// Delimiters
	COMMA     = "COMMA"
	SEMICOLON = "SEMICOLON"

	LPAREN = "LPAREN"
	RPAREN = "RPAREN"
	LBRACE = "LBRACE"
	RBRACE = "RBRACE"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(t TokenType, char byte) Token {
	return Token{Type: t, Literal: string(char)}
}

func IsTokenEq(t1, t2 Token) bool {
	return t1.Type == t2.Type && t1.Literal == t2.Literal
}
