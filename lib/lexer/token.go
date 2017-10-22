package lexer

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Literals
	INT = "INT" // 1, 2, 333

	// Variable names
	IDENT = "IDENT"

	// Operators
	ASSIGN   = "ASSIGN"
	PLUS     = "PLUS"
	MINUS    = "MINUS"
	BANG     = "BANG"
	ASTERISK = "ASTERISK"
	SLASH    = "SLASH"
	LT       = "LT"
	GT       = "GT"
	EQ       = "EQ"
	NOT_EQ   = "NOT_EQ"

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
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// IdentType of an identifier is either one of the keywords or IDENT
func IdentType(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

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
