package parser_test

import (
	"fmt"
	"testing"

	"github.com/avazquezcode/govetryx/internal/usecase/parser"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	tests := map[string]struct {
		parsingErr  parser.ParsingErr
		expectedErr string
	}{
		"happy path (errors slice not empty)": {
			parsingErr: *parser.NewParsingErr([]error{
				fmt.Errorf("error A"),
				fmt.Errorf("error B"),
				fmt.Errorf("error C"),
			}),
			expectedErr: "error #1: error A \nerror #2: error B \nerror #3: error C \n",
		},
		"empty slice)": {
			parsingErr:  *parser.NewParsingErr([]error{}),
			expectedErr: "",
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			assert.Equal(t, test.expectedErr, test.parsingErr.Error())
		})
	}
}
