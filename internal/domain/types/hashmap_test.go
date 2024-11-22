package types_test

import (
	"testing"

	"github.com/avazquezcode/govetryx/internal/domain/types"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	tests := map[string]struct {
		hashMap  *types.HashMap
		key      interface{}
		expected interface{}
	}{
		"valid key": {
			hashMap: &types.HashMap{
				"key": "value",
			},
			key:      "key",
			expected: "value",
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			assert.Equal(t, test.expected, test.hashMap.Get(test.key))
		})
	}
}

func TestExists(t *testing.T) {
	tests := map[string]struct {
		hashMap  *types.HashMap
		key      interface{}
		expected bool
	}{
		"valid key": {
			hashMap: &types.HashMap{
				"key": "value",
			},
			key:      "key",
			expected: true,
		},
		"key not found": {
			hashMap: &types.HashMap{
				"key": "value",
			},
			key:      "anything",
			expected: false,
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			assert.Equal(t, test.expected, test.hashMap.Exists(test.key))
		})
	}
}

func TestSet(t *testing.T) {
	tests := map[string]struct {
		hashMap  *types.HashMap
		key      interface{}
		value    interface{}
		expected *types.HashMap
	}{
		"valid key": {
			hashMap: &types.HashMap{},
			key:     "key",
			value:   "value",
			expected: &types.HashMap{
				"key": "value",
			},
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			test.hashMap.Set(test.key, test.value)
			assert.Equal(t, test.expected, test.hashMap)
		})
	}
}
