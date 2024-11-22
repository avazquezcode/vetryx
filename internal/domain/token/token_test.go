package token_test

import (
	"testing"

	"github.com/avazquezcode/govetryx/internal/domain/token"
	"github.com/stretchr/testify/assert"
)

func TestNewToken(t *testing.T) {
	tests := map[string]struct {
		tokenType token.Type
		lexeme    string
		literal   interface{}
		line      int
		expected  *token.Token
	}{
		"valid construction": {
			tokenType: token.And,
			lexeme:    "&",
			literal:   nil,
			line:      2, // could be any line
			expected: &token.Token{
				Type:    token.And,
				Lexeme:  "&",
				Literal: nil,
				Line:    2,
			},
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			token := token.NewToken(test.tokenType, test.lexeme, test.literal, test.line)
			assert.Equal(t, test.expected, token)
		})
	}
}
