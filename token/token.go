package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EPF"
	IDENT   = "IDENT"
	INT     = "INT"
	ASSIGN = "="
	PLUS = "+"
	COMMA = ","
	SEMICOLON=";"
	LPAREN="("
	RPAREN=")"
	LBRACE="{"
	RBRACE="}"

	FUNCTION="FUNCTION"
	LET="LET"
)
