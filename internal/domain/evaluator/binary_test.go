package evaluator_test

import (
	"testing"

	"github.com/avazquezcode/govetryx/internal/domain/evaluator"
	"github.com/avazquezcode/govetryx/internal/domain/token"

	"github.com/stretchr/testify/assert"
)

func TestDifferent(t *testing.T) {
	tests := map[string]struct {
		left    interface{}
		right   interface{}
		want    bool
		wantErr bool
	}{
		"equal values": {
			left:    1,
			right:   1,
			want:    false,
			wantErr: false,
		},
		"different values": {
			left:    1,
			right:   2,
			want:    true,
			wantErr: false,
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			evaluator, _ := evaluator.NewBinaryEvaluator(test.left, token.NewToken(token.NotEqual, "!=", "", 1), test.right)
			result, err := evaluator.Evaluate()
			assert.Equal(t, test.want, result)
			assert.Equal(t, test.wantErr, err != nil)
		})
	}
}

func TestEqual(t *testing.T) {
	tests := map[string]struct {
		left    interface{}
		right   interface{}
		want    bool
		wantErr bool
	}{
		"equal values": {
			left:    1,
			right:   1,
			want:    true,
			wantErr: false,
		},
		"different values": {
			left:    1,
			right:   2,
			want:    false,
			wantErr: false,
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			evaluator, _ := evaluator.NewBinaryEvaluator(test.left, token.NewToken(token.EqualEqual, "==", "", 1), test.right)
			result, err := evaluator.Evaluate()
			assert.Equal(t, test.want, result)
			assert.Equal(t, test.wantErr, err != nil)
		})
	}
}

func TestGreaterOrEqual(t *testing.T) {
	tests := map[string]struct {
		left    interface{}
		right   interface{}
		want    interface{}
		wantErr bool
	}{
		"left equal to right": {
			left:    float64(1),
			right:   float64(1),
			want:    true,
			wantErr: false,
		},
		"left lower than right": {
			left:    float64(1),
			right:   float64(2),
			want:    false,
			wantErr: false,
		},
		"left greater than right": {
			left:    float64(3),
			right:   float64(2),
			want:    true,
			wantErr: false,
		},
		"error: one of the values is not a float": {
			left:    3,
			right:   float64(2),
			want:    nil,
			wantErr: true,
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			evaluator, _ := evaluator.NewBinaryEvaluator(test.left, token.NewToken(token.GreaterOrEqual, ">=", "", 1), test.right)
			result, err := evaluator.Evaluate()
			assert.Equal(t, test.want, result)
			assert.Equal(t, test.wantErr, err != nil)
		})
	}
}

func TestGreater(t *testing.T) {
	tests := map[string]struct {
		left    interface{}
		right   interface{}
		want    interface{}
		wantErr bool
	}{
		"left equal to right": {
			left:    float64(1),
			right:   float64(1),
			want:    false,
			wantErr: false,
		},
		"left lower than right": {
			left:    float64(1),
			right:   float64(2),
			want:    false,
			wantErr: false,
		},
		"left greater than right": {
			left:    float64(3),
			right:   float64(2),
			want:    true,
			wantErr: false,
		},
		"error: one of the values is not a float": {
			left:    3,
			right:   float64(2),
			want:    nil,
			wantErr: true,
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			evaluator, _ := evaluator.NewBinaryEvaluator(test.left, token.NewToken(token.Greater, ">", "", 1), test.right)
			result, err := evaluator.Evaluate()
			assert.Equal(t, test.want, result)
			assert.Equal(t, test.wantErr, err != nil)
		})
	}
}

func TestLowerOrEqual(t *testing.T) {
	tests := map[string]struct {
		left    interface{}
		right   interface{}
		want    interface{}
		wantErr bool
	}{
		"left equal to right": {
			left:    float64(1),
			right:   float64(1),
			want:    true,
			wantErr: false,
		},
		"left lower than right": {
			left:    float64(1),
			right:   float64(2),
			want:    true,
			wantErr: false,
		},
		"left greater than right": {
			left:    float64(3),
			right:   float64(2),
			want:    false,
			wantErr: false,
		},
		"error: one of the values is not a float": {
			left:    3,
			right:   float64(2),
			want:    nil,
			wantErr: true,
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			evaluator, _ := evaluator.NewBinaryEvaluator(test.left, token.NewToken(token.LowerOrEqual, "<=", "", 1), test.right)
			result, err := evaluator.Evaluate()
			assert.Equal(t, test.want, result)
			assert.Equal(t, test.wantErr, err != nil)
		})
	}
}

func TestLower(t *testing.T) {
	tests := map[string]struct {
		left    interface{}
		right   interface{}
		want    interface{}
		wantErr bool
	}{
		"left equal to right": {
			left:    float64(1),
			right:   float64(1),
			want:    false,
			wantErr: false,
		},
		"left lower than right": {
			left:    float64(1),
			right:   float64(2),
			want:    true,
			wantErr: false,
		},
		"left greater than right": {
			left:    float64(3),
			right:   float64(2),
			want:    false,
			wantErr: false,
		},
		"error: one of the values is not a float": {
			left:    3,
			right:   float64(2),
			want:    nil,
			wantErr: true,
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			evaluator, _ := evaluator.NewBinaryEvaluator(test.left, token.NewToken(token.Lower, "<", "", 1), test.right)
			result, err := evaluator.Evaluate()
			assert.Equal(t, test.want, result)
			assert.Equal(t, test.wantErr, err != nil)
		})
	}
}

func TestMinus(t *testing.T) {
	tests := map[string]struct {
		left    interface{}
		right   interface{}
		want    interface{}
		wantErr bool
	}{
		"subtraction that returns a positive value": {
			left:    float64(2),
			right:   float64(1),
			want:    float64(1),
			wantErr: false,
		},
		"subtraction that returns a positive value with decimal": {
			left:    float64(2.5),
			right:   float64(1),
			want:    float64(1.5),
			wantErr: false,
		},
		"subtraction that returns zero": {
			left:    float64(2.5),
			right:   float64(2.5),
			want:    float64(0),
			wantErr: false,
		},
		"subtraction that returns a negative value": {
			left:    float64(1),
			right:   float64(2),
			want:    float64(-1),
			wantErr: false,
		},
		"subtraction that returns a negative value with decimal": {
			left:    float64(1),
			right:   float64(2.5),
			want:    float64(-1.5),
			wantErr: false,
		},
		"subtraction with input that is not valid": {
			left:    "text",
			right:   float64(2.5),
			want:    nil,
			wantErr: true,
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			evaluator, _ := evaluator.NewBinaryEvaluator(test.left, token.NewToken(token.Minus, "-", "", 1), test.right)
			result, err := evaluator.Evaluate()
			assert.Equal(t, test.want, result)
			assert.Equal(t, test.wantErr, err != nil)
		})
	}
}

func TestPlus(t *testing.T) {
	tests := map[string]struct {
		left    interface{}
		right   interface{}
		want    interface{}
		wantErr bool
	}{
		"addition that returns a positive value": {
			left:    float64(2),
			right:   float64(1),
			want:    float64(3),
			wantErr: false,
		},
		"addition that returns a positive value with decimal": {
			left:    float64(2.5),
			right:   float64(1),
			want:    float64(3.5),
			wantErr: false,
		},
		"addition that returns zero": {
			left:    float64(2.5),
			right:   float64(-2.5),
			want:    float64(0),
			wantErr: false,
		},
		"addition that returns a negative value": {
			left:    float64(-1),
			right:   float64(-2),
			want:    float64(-3),
			wantErr: false,
		},
		"addition that returns a negative value with decimal": {
			left:    float64(1),
			right:   float64(-2.5),
			want:    float64(-1.5),
			wantErr: false,
		},
		"addition with input that is not valid": {
			left:    "text",
			right:   float64(2.5),
			want:    nil,
			wantErr: true,
		},
		"concatenation (case of left and right are strings)": {
			left:    "hello ",
			right:   "world",
			want:    "hello world",
			wantErr: false,
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			evaluator, _ := evaluator.NewBinaryEvaluator(test.left, token.NewToken(token.Plus, "+", "", 1), test.right)
			result, err := evaluator.Evaluate()
			assert.Equal(t, test.want, result)
			assert.Equal(t, test.wantErr, err != nil)
		})
	}
}

func TestSlash(t *testing.T) {
	tests := map[string]struct {
		left    interface{}
		right   interface{}
		want    interface{}
		wantErr bool
	}{
		"division that returns a positive value": {
			left:    float64(4),
			right:   float64(2),
			want:    float64(2),
			wantErr: false,
		},
		"division that returns a positive value with decimal": {
			left:    float64(2.5),
			right:   float64(1),
			want:    float64(2.5),
			wantErr: false,
		},
		"division per zero - should throw error": {
			left:    float64(2.5),
			right:   float64(0),
			want:    nil,
			wantErr: true,
		},
		"left or right invalid": {
			left:    "invalid",
			right:   float64(2),
			want:    nil,
			wantErr: true,
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			evaluator, _ := evaluator.NewBinaryEvaluator(test.left, token.NewToken(token.Slash, "/", "", 1), test.right)
			result, err := evaluator.Evaluate()
			assert.Equal(t, test.want, result)
			assert.Equal(t, test.wantErr, err != nil)
		})
	}
}

func TestStar(t *testing.T) {
	tests := map[string]struct {
		left    interface{}
		right   interface{}
		want    interface{}
		wantErr bool
	}{
		"multiplication that returns a positive value": {
			left:    float64(4),
			right:   float64(2),
			want:    float64(8),
			wantErr: false,
		},
		"multiplication that returns a positive value with decimal": {
			left:    float64(2.5),
			right:   float64(1),
			want:    float64(2.5),
			wantErr: false,
		},
		"multiplication per zero - return 0": {
			left:    float64(2.5),
			right:   float64(0),
			want:    float64(0),
			wantErr: false,
		},
		"left or right invalid": {
			left:    "invalid",
			right:   float64(2),
			want:    nil,
			wantErr: true,
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			evaluator, _ := evaluator.NewBinaryEvaluator(test.left, token.NewToken(token.Star, "*", "", 1), test.right)
			result, err := evaluator.Evaluate()
			assert.Equal(t, test.want, result)
			assert.Equal(t, test.wantErr, err != nil)
		})
	}
}

func TestModulus(t *testing.T) {
	tests := map[string]struct {
		left    interface{}
		right   interface{}
		want    interface{}
		wantErr bool
	}{
		"modulus that returns zero": {
			left:    float64(4),
			right:   float64(2),
			want:    float64(0),
			wantErr: false,
		},
		"modulus that returns positive": {
			left:    float64(4),
			right:   float64(3),
			want:    float64(1),
			wantErr: false,
		},
		"modulus that returns negative": {
			left:    float64(-4),
			right:   float64(3),
			want:    float64(-1),
			wantErr: false,
		},
		"modulus with right operand being 0": {
			left:    float64(-4),
			right:   float64(0),
			want:    nil,
			wantErr: true,
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			evaluator, _ := evaluator.NewBinaryEvaluator(test.left, token.NewToken(token.Modulus, "%", "", 1), test.right)
			result, err := evaluator.Evaluate()
			assert.Equal(t, test.want, result)
			assert.Equal(t, test.wantErr, err != nil)
		})
	}
}
