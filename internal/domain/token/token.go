// Token is a core package of our interpreter, since it defined the different accepted tokens in the language.
package token

// Token represents a token of the language.
type Token struct {
	Type    Type
	Lexeme  string
	Literal interface{}
	Line    int
}

// NewToken is a constructor for a new token.
func NewToken(tokenType Type, lexeme string, literal interface{}, line int) *Token {
	return &Token{
		Type:    tokenType,
		Lexeme:  lexeme,
		Literal: literal,
		Line:    line,
	}
}
