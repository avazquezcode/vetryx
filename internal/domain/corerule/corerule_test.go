package corerule_test

import (
	"testing"

	"github.com/avazquezcode/govetryx/internal/domain/corerule"
	"github.com/stretchr/testify/assert"
)

func TestIsTrue(t *testing.T) {
	tests := map[string]struct {
		value    interface{}
		expected bool
	}{
		"input is a TRUE boolean": {
			value:    true,
			expected: true,
		},
		"input is a FALSE boolean": {
			value:    false,
			expected: false,
		},
		"input is null": {
			value:    nil,
			expected: false,
		},
		"input is something else than nil or boolean": {
			value:    "1",
			expected: true,
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			assert.Equal(t, test.expected, corerule.IsTrue(test.value))
		})
	}
}

func TestIsEqual(t *testing.T) {
	tests := map[string]struct {
		a        interface{}
		b        interface{}
		expected bool
	}{
		"true = true": {
			a:        true,
			b:        true,
			expected: true,
		},
		"true != false": {
			a:        true,
			b:        false,
			expected: false,
		},
		"nil = nil": {
			a:        nil,
			b:        nil,
			expected: true,
		},
		"string = string": {
			a:        "a",
			b:        "a",
			expected: true,
		},
		"string <> string": {
			a:        "a",
			b:        "b",
			expected: false,
		},
		"number = number": {
			a:        "1",
			b:        "1",
			expected: true,
		},
		"number <> number": {
			a:        "1",
			b:        "2",
			expected: false,
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			assert.Equal(t, test.expected, corerule.IsEqual(test.a, test.b))
		})
	}
}

func TestPrintableValue(t *testing.T) {
	tests := map[string]struct {
		value    interface{}
		expected string
	}{
		"true bool": {
			value:    true,
			expected: "true",
		},
		"false bool": {
			value:    false,
			expected: "false",
		},
		"null": {
			value:    nil,
			expected: "null",
		},
		"empty string": {
			value:    "",
			expected: "",
		},
		"valid string": {
			value:    "asd",
			expected: "asd",
		},
		"valid number": {
			value:    1,
			expected: "1",
		},
		"valid number = 0": {
			value:    0,
			expected: "0",
		},
		"valid number = 1.751": {
			value:    1.751,
			expected: "1.751",
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			assert.Equal(t, test.expected, corerule.PrintableValue(test.value))
		})
	}
}
