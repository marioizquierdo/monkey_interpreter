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
	l.skipWhitespace()

	var tok Token
	c := l.char
	switch c {
	case '=':
		tok = NewToken(ASSIGN, c)
	case '+':
		tok = NewToken(PLUS, c)
	case ',':
		tok = NewToken(COMMA, c)
	case ';':
		tok = NewToken(SEMICOLON, c)
	case '(':
		tok = NewToken(LPAREN, c)
	case ')':
		tok = NewToken(RPAREN, c)
	case '{':
		tok = NewToken(LBRACE, c)
	case '}':
		tok = NewToken(RBRACE, c)
	case 0:
		tok = Token{Type: EOF, Literal: ""}
	default:
		if isLetter(c) {
			tok.Literal = l.readIdentifier()
			tok.Type = IdentType(tok.Literal) // IDENT or a reserved keyword
			return tok
		} else if isDigit(c) {
			tok.Type = INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = NewToken(ILLEGAL, c)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	pos := l.pos
	for isLetter(l.char) {
		l.readChar()
	}
	return l.input[pos:l.pos]
}

func (l *Lexer) readNumber() string {
	pos := l.pos
	for isDigit(l.char) {
		l.readChar()
	}
	return l.input[pos:l.pos]
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || // lowercase
		'A' <= char && char <= 'Z' || // uppercase
		char == '_' // separator
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}
