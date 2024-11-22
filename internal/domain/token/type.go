package token

// Type represents a token type.
type Type int

const (
	// Loop
	While Type = iota
	Break
	Continue

	// Conditions
	If
	Else

	// Other operators
	And
	Or

	// Functions
	Fn
	Return

	// Types
	Identifier
	Number
	String
	False
	True
	Null

	// Variables
	VarDeclarator
	VarShortDeclarator

	// 1 character tokens
	Bang
	Plus
	Minus
	Modulus
	LeftBrace
	RightBrace
	LeftParentheses
	RightParentheses
	Comma
	Slash
	Hashtag
	Star
	Equal
	EqualEqual
	NotEqual
	Lower
	LowerOrEqual
	Greater
	GreaterOrEqual
	Semicolon

	// Inbuilt functions
	Print

	// End of file
	EOF
)

// ReservedWordsMapper is a map of our reserved words.
var ReservedWordsMapper = map[string]Type{
	"dec":      VarDeclarator,
	"fn":       Fn,
	"true":     True,
	"false":    False,
	"if":       If,
	"else":     Else,
	"while":    While,
	"break":    Break,
	"continue": Continue,
	"print":    Print,
	"return":   Return,
	"null":     Null,
}
