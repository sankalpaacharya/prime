package lexer

import (
	"prime/token"
	"testing"
)

func TestNextToken(t *testing.T) {

	input := `=+(){},;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRAC, "{"},
		{token.RBRAC, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := New(input)

	for i, tt := range tests {

		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] tokentype wrong. expected=%q got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d] expected literal wrong. expected=%q got=%q", i, tt.expectedLiteral, tok.Literal)
		}

	}

}
