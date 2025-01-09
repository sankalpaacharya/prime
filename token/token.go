package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Idenfier and Lietral
	IDENT = "IDENT"
	INT   = "INT"

	// operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	DIVISION = "/"

	LPAREN = "("
	RPAREN = ")"
	LBRAC  = "{"
	RBRAC  = "}"

	// delimeters
	COMMA     = ","
	SEMICOLON = ";"

	// Keywords
	FUNCTION = "function"
	LET      = "let"
)
