package evaluator_test

import (
	"testing"

	"github.com/avazquezcode/govetryx/internal/domain/evaluator"
	"github.com/avazquezcode/govetryx/internal/domain/token"

	"github.com/stretchr/testify/assert"
)

func TestBangNegation(t *testing.T) {
	tests := map[string]struct {
		expression interface{}
		want       interface{}
		wantErr    bool
	}{
		"bang on false": {
			expression: true,
			want:       false,
			wantErr:    false,
		},
		"bang on true": {
			expression: false,
			want:       true,
			wantErr:    false,
		},
		"bang on nil": {
			expression: nil,
			want:       true,
			wantErr:    false,
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			evaluator, _ := evaluator.NewUnaryEvaluator(token.NewToken(token.Bang, "!", "", 1), test.expression)
			result, err := evaluator.Evaluate()
			assert.Equal(t, test.want, result)
			assert.Equal(t, test.wantErr, err != nil)
		})
	}
}

func TestMinusNegation(t *testing.T) {
	tests := map[string]struct {
		expression interface{}
		want       interface{}
		wantErr    bool
	}{
		"negation of positive": {
			expression: float64(1),
			want:       float64(-1),
			wantErr:    false,
		},
		"negation of negative": {
			expression: float64(-1),
			want:       float64(1),
			wantErr:    false,
		},
		"not a number": {
			expression: "hey",
			want:       nil,
			wantErr:    true,
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			evaluator, _ := evaluator.NewUnaryEvaluator(token.NewToken(token.Minus, "-", "", 1), test.expression)
			result, err := evaluator.Evaluate()
			assert.Equal(t, test.want, result)
			assert.Equal(t, test.wantErr, err != nil)
		})
	}
}
