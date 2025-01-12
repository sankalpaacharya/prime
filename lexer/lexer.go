package lexer

import (
	"prime/token"
)

type Lexer struct {
	input        string
	position     int
	nextPosition int
	char         byte
}

// read a character from input and move the pointer forward
func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.nextPosition]
	}
	l.position = l.nextPosition
	l.nextPosition += 1
}

// just to initialize Lexer struct
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhiteSpace()

	switch l.char {

	case '=':
		tok = newToken(token.ASSIGN, '=')

	case '+':
		tok = newToken(token.PLUS, '+')

	case '-':
		tok = newToken(token.MINUS, '-')

	case '*':
		tok = newToken(token.ASTERISK, '*')

	case '/':
		tok = newToken(token.SLASH, '/')

	case ')':
		tok = newToken(token.RPAREN, ')')

	case '(':
		tok = newToken(token.LPAREN, '(')

	case '}':
		tok = newToken(token.RBRAC, '}')

	case '{':
		tok = newToken(token.LBRAC, '{')

	case ';':
		tok = newToken(token.SEMICOLON, ';')

	case ',':
		tok = newToken(token.COMMA, ',')
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.char) {
			tok.Literal = l.readIdentifier()
			tok.Type = LookUpIdent(tok.Literal)
			return tok
		} else if isDigit(l.char) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.char)
		}
	}

	l.readChar()
	return tok
}

// read the identifier
func (l *Lexer) readIdentifier() string {
	pos := l.position
	for isLetter(l.char) {
		l.readChar()
	}
	return l.input[pos:l.position]
}
func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
func (l *Lexer) readNumber() string {
	pos := l.position
	for isDigit(l.char) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) skipWhiteSpace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

// check if it a character
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && 'Z' >= ch || ch == '_'
}

// make a token
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func LookUpIdent(iden string) token.TokenType {
	if ident, ok := token.Keywords[iden]; ok {
		return ident
	}
	return token.IDENT
}
