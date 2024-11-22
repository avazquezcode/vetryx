package evaluator

import (
	"fmt"

	"github.com/avazquezcode/govetryx/internal/domain/corerule"
)

type (
	// BangNegation evaluates the negation of something. (Eg: !true => false).
	BangNegation struct {
		expression interface{}
	}

	// MinusNegation evaluates the negative value of a number. (Eg: -(-1) => 1).
	MinusNegation struct {
		expression interface{}
	}
)

func (e *BangNegation) Evaluate() (interface{}, error) {
	return !corerule.IsTrue(e.expression), nil
}

func (e *MinusNegation) Evaluate() (interface{}, error) {
	if isNumber(e.expression) {
		return -e.expression.(float64), nil
	}
	return nil, fmt.Errorf("not a number")
}
