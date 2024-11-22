package scanner

import (
	"bytes"
	"testing"

	"github.com/avazquezcode/govetryx/internal/domain/token"
	"github.com/stretchr/testify/assert"
)

func TestScan(t *testing.T) {
	tests := map[string]struct {
		src         string
		expected    []*token.Token
		expectedErr bool
	}{
		"empty src": {
			src: "",
			expected: []*token.Token{
				token.NewToken(token.EOF, "", nil, 1),
			},
		},
		"ignored characters": {
			src: " \t \t \r    ",
			expected: []*token.Token{
				token.NewToken(token.EOF, "", nil, 1),
			},
		},
		"new lines": {
			src: "\n \n \n",
			expected: []*token.Token{
				token.NewToken(token.EOF, "", nil, 4),
			},
		},
		"braces": {
			src: "{}",
			expected: []*token.Token{
				token.NewToken(token.LeftBrace, "{", nil, 1),
				token.NewToken(token.RightBrace, "}", nil, 1),
				token.NewToken(token.EOF, "", nil, 1),
			},
		},
		"parentheses": {
			src: "(())",
			expected: []*token.Token{
				token.NewToken(token.LeftParentheses, "(", nil, 1),
				token.NewToken(token.LeftParentheses, "(", nil, 1),
				token.NewToken(token.RightParentheses, ")", nil, 1),
				token.NewToken(token.RightParentheses, ")", nil, 1),
				token.NewToken(token.EOF, "", nil, 1),
			},
		},
		"arithmetic operators": {
			src: "+-*/%",
			expected: []*token.Token{
				token.NewToken(token.Plus, "+", nil, 1),
				token.NewToken(token.Minus, "-", nil, 1),
				token.NewToken(token.Star, "*", nil, 1),
				token.NewToken(token.Slash, "/", nil, 1),
				token.NewToken(token.Modulus, "%", nil, 1),
				token.NewToken(token.EOF, "", nil, 1),
			},
		},
		"comma": {
			src: ",",
			expected: []*token.Token{
				token.NewToken(token.Comma, ",", nil, 1),
				token.NewToken(token.EOF, "", nil, 1),
			},
		},
		"equal (assignment)": {
			src: "=",
			expected: []*token.Token{
				token.NewToken(token.Equal, "=", nil, 1),
				token.NewToken(token.EOF, "", nil, 1),
			},
		},
		"double equal (comparison)": {
			src: "==",
			expected: []*token.Token{
				token.NewToken(token.EqualEqual, "==", nil, 1),
				token.NewToken(token.EOF, "", nil, 1),
			},
		},
		"greater": {
			src: ">",
			expected: []*token.Token{
				token.NewToken(token.Greater, ">", nil, 1),
				token.NewToken(token.EOF, "", nil, 1),
			},
		},
		"greater or equal": {
			src: ">=",
			expected: []*token.Token{
				token.NewToken(token.GreaterOrEqual, ">=", nil, 1),
				token.NewToken(token.EOF, "", nil, 1),
			},
		},
		"lower": {
			src: "<",
			expected: []*token.Token{
				token.NewToken(token.Lower, "<", nil, 1),
				token.NewToken(token.EOF, "", nil, 1),
			},
		},
		"lower or equal": {
			src: "<=",
			expected: []*token.Token{
				token.NewToken(token.LowerOrEqual, "<=", nil, 1),
				token.NewToken(token.EOF, "", nil, 1),
			},
		},
		"not equal": {
			src: "<>",
			expected: []*token.Token{
				token.NewToken(token.NotEqual, "<>", nil, 1),
				token.NewToken(token.EOF, "", nil, 1),
			},
		},
		"bang operator": {
			src: "!",
			expected: []*token.Token{
				token.NewToken(token.Bang, "!", nil, 1),
				token.NewToken(token.EOF, "", nil, 1),
			},
		},

		"short variable declarator": {
			src: ":=",
			expected: []*token.Token{
				token.NewToken(token.VarShortDeclarator, ":=", nil, 1),
				token.NewToken(token.EOF, "", nil, 1),
			},
		},
		"and": {
			src: "&&",
			expected: []*token.Token{
				token.NewToken(token.And, "&&", nil, 1),
				token.NewToken(token.EOF, "", nil, 1),
			},
		},
		"or": {
			src: "||",
			expected: []*token.Token{
				token.NewToken(token.Or, "||", nil, 1),
				token.NewToken(token.EOF, "", nil, 1),
			},
		},
		"# comments (everything after is ignored)": {
			src: "# anything1234",
			expected: []*token.Token{
				token.NewToken(token.EOF, "", nil, 1),
			},
		},
		"empty string": {
			src: `""`,
			expected: []*token.Token{
				token.NewToken(token.String, "\"\"", "", 1),
				token.NewToken(token.EOF, "", nil, 1),
			},
		},
		"non empty string": {
			src: `"hello"`,
			expected: []*token.Token{
				token.NewToken(token.String, "\"hello\"", "hello", 1),
				token.NewToken(token.EOF, "", nil, 1),
			},
		},
		"number (but as string)": {
			src: `"123"`,
			expected: []*token.Token{
				token.NewToken(token.String, "\"123\"", "123", 1),
				token.NewToken(token.EOF, "", nil, 1),
			},
		},
		"number": {
			src: `123`,
			expected: []*token.Token{
				token.NewToken(token.Number, "123", float64(123), 1),
				token.NewToken(token.EOF, "", nil, 1),
			},
		},
		"number with decimals": {
			src: `123.12`,
			expected: []*token.Token{
				token.NewToken(token.Number, "123.12", float64(123.12), 1),
				token.NewToken(token.EOF, "", nil, 1),
			},
		},
		"0": {
			src: `0`,
			expected: []*token.Token{
				token.NewToken(token.Number, "0", float64(0), 1),
				token.NewToken(token.EOF, "", nil, 1),
			},
		},
		"identifier": {
			src: "hi",
			expected: []*token.Token{
				token.NewToken(token.Identifier, "hi", nil, 1),
				token.NewToken(token.EOF, "", nil, 1),
			},
		},
		"reserved words": {
			src: `dec fn true false if else while print return null break continue`,
			expected: []*token.Token{
				token.NewToken(token.VarDeclarator, "dec", nil, 1),
				token.NewToken(token.Fn, "fn", nil, 1),
				token.NewToken(token.True, "true", nil, 1),
				token.NewToken(token.False, "false", nil, 1),
				token.NewToken(token.If, "if", nil, 1),
				token.NewToken(token.Else, "else", nil, 1),
				token.NewToken(token.While, "while", nil, 1),
				token.NewToken(token.Print, "print", nil, 1),
				token.NewToken(token.Return, "return", nil, 1),
				token.NewToken(token.Null, "null", nil, 1),
				token.NewToken(token.Break, "break", nil, 1),
				token.NewToken(token.Continue, "continue", nil, 1),
				token.NewToken(token.EOF, "", nil, 1),
			},
		},
		"unknown character": {
			src:         `?`,
			expectedErr: true,
		},
		"string that starts but doesn't end": {
			src:         `"`,
			expectedErr: true,
		},
		"invalid float": {
			src:         `123...`,
			expectedErr: true,
		},
		"invalid float - another case": {
			src:         `123.a1s`,
			expectedErr: true,
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			scanner := NewScanner(bytes.Runes(strToBytes(test.src)))
			tokens, err := scanner.Scan()

			if test.expectedErr {
				assert.NotNil(t, err)
				return
			}

			assert.Equal(t, test.expected, tokens)
			assert.Nil(t, err)
		})
	}
}

func strToBytes(str string) []byte {
	return []byte(str)
}
