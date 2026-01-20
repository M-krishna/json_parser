package tokens

type TokenType int

const (
	// Structural tokens
	TOKEN_OPEN_BRACE TokenType = iota	// {
	TOKEN_CLOSE_BRACE					// }
	TOKEN_OPEN_BRACKET					// [
	TOKEN_CLOSE_BRACKET					// ]
	TOKEN_COLON							// :
	TOKEN_COMMA							// ,

	// Value tokens
	TOKEN_STRING
	TOKEN_NUMBER

	// Literals
	TOKEN_TRUE
	TOKEN_FALSE
	TOKEN_NULL

	// End of input
	TOKEN_EOF

)

func (t TokenType) String() string {
	switch t {
	case TOKEN_OPEN_BRACE:
		return "TOKEN_OPEN_BRACE"
	case TOKEN_CLOSE_BRACE:
		return "TOKEN_CLOSE_BRACE"
	case TOKEN_OPEN_BRACKET:
		return "TOKEN_OPEN_BRACKET"
	case TOKEN_CLOSE_BRACKET:
		return "TOKEN_CLOSE_BRACKET"
	case TOKEN_COLON:
		return "TOKEN_COLON"
	case TOKEN_COMMA:
		return "TOKEN_COMMA"
	case TOKEN_EOF:
		return "TOKEN_EOF"
	}
	return ""
}

type Token struct {
	Type 	TokenType
	Start 	int		// index into the byte slice (inclusive)
	End		int		// index into the byte slice (exclusive)
}

func NewToken(tokenType TokenType, start int, end int) Token {
	return Token{
		Type: tokenType,
		Start: start,
		End: end,
	}
}