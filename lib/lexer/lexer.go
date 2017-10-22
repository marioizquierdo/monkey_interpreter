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

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	var tok Token
	c := l.char
	switch c {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok = Token{Type: EQ, Literal: "=="}
		} else {
			tok = NewToken(ASSIGN, c)
		}
	case '+':
		tok = NewToken(PLUS, c)
	case '-':
		tok = NewToken(MINUS, c)
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = Token{Type: NOT_EQ, Literal: "!="}
		} else {
			tok = NewToken(BANG, c)
		}
	case '*':
		tok = NewToken(ASTERISK, c)
	case '/':
		tok = NewToken(SLASH, c)
	case '<':
		tok = NewToken(LT, c)
	case '>':
		tok = NewToken(GT, c)
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

func (l *Lexer) peekChar() byte {
	if l.readPos >= l.inputLen {
		return 0
	}
	return l.input[l.readPos]
}

func (l *Lexer) readChar() {
	l.char = l.peekChar()
	l.pos = l.readPos
	l.readPos += 1
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
