package types_test

import (
	"testing"

	"github.com/avazquezcode/govetryx/internal/domain/types"
	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	tests := map[string]struct {
		expected types.Stack
	}{
		"valid construction": {
			expected: nil,
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			assert.Equal(t, test.expected, types.NewStack())
		})
	}
}
func TestPush(t *testing.T) {
	tests := map[string]struct {
		stack    *types.Stack
		element  interface{}
		expected *types.Stack
	}{
		"add element to the stack": {
			stack:   &types.Stack{},
			element: "test",
			expected: &types.Stack{
				"test",
			},
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			test.stack.Push(test.element)
			assert.Equal(t, test.expected, test.stack)
		})
	}
}

func TestPeek(t *testing.T) {
	tests := map[string]struct {
		stack    *types.Stack
		expected interface{}
	}{
		"valid element on top of stack": {
			stack: &types.Stack{
				"test", "abc",
			},
			expected: "abc",
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			assert.Equal(t, test.expected, test.stack.Peek())
		})
	}
}

func TestPop(t *testing.T) {
	tests := map[string]struct {
		stack    *types.Stack
		expected interface{}
	}{
		"pop from empty stack": {
			stack:    &types.Stack{},
			expected: &types.Stack{},
		},
		"pop from stack with 1 value": {
			stack: &types.Stack{
				"test",
			},
			expected: &types.Stack{},
		},
		"pop from stack with 2 values": {
			stack: &types.Stack{
				"test", "abc",
			},
			expected: &types.Stack{
				"test",
			},
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			test.stack.Pop()
			assert.Equal(t, test.expected, test.stack)
		})
	}
}

func TestLength(t *testing.T) {
	tests := map[string]struct {
		stack    *types.Stack
		expected int
	}{
		"empty": {
			stack:    &types.Stack{},
			expected: 0,
		},
		"1 element": {
			stack:    &types.Stack{"test"},
			expected: 1,
		},
		"2 elements": {
			stack:    &types.Stack{"test", "abc"},
			expected: 2,
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			assert.Equal(t, test.expected, test.stack.Length())
		})
	}
}
