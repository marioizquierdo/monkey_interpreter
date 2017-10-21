package lexer

type Lexer struct {
	input    string // input program to be tokenized
	inputLen int    // number of chars on input

	pos     int  // index to current char
	readPos int  // index to next char, used when peeking into multi-char tokens
	char    byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input, inputLen: len(input)}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPos < l.inputLen {
		l.char = l.input[l.readPos]
	} else {
		l.char = 0 // NUL, end of file
	}
	l.pos = l.readPos
	l.readPos += 1
}

func (l *Lexer) NextToken() Token {
	var tok Token

	switch l.char {
	case '=':
		tok = NewToken(ASSIGN, l.char)
	case '+':
		tok = NewToken(PLUS, l.char)
	case ',':
		tok = NewToken(COMMA, l.char)
	case ';':
		tok = NewToken(SEMICOLON, l.char)
	case '(':
		tok = NewToken(LPAREN, l.char)
	case ')':
		tok = NewToken(RPAREN, l.char)
	case '{':
		tok = NewToken(LBRACE, l.char)
	case '}':
		tok = NewToken(RBRACE, l.char)
	case 0:
		tok = Token{Type: EOF, Literal: ""}
	}

	l.readChar()
	return tok
}
