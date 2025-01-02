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

func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.position]
	}
	l.position = l.nextPosition
	l.nextPosition += 1
}

func Next(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.char {

	case '=':
		tok = newToken(token.ASSIGN, '=')

	case '+':
		tok = newToken(token.PLUS, '+')

	case '-':
		tok = newToken(token.MINUS, '-')

	case '*':
		tok = newToken(token.MULTIPLY, '*')

	case '/':
		tok = newToken(token.DIVISION, '/')

	case ')':
		tok = newToken(token.LPAREN, ')')

	case '(':
		tok = newToken(token.RPAREN, '(')

	case '}':
		tok = newToken(token.LBRAC, '=')

	case '{':
		tok = newToken(token.RBRAC, '{')

	case ';':
		tok = newToken(token.SEMICOLON, ';')

	case ',':
		tok = newToken(token.COMMA, ',')

	}
	return tok

}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}

}
