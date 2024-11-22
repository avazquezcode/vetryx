package evaluator

import (
	interr "github.com/avazquezcode/govetryx/internal/domain/error"
	"github.com/avazquezcode/govetryx/internal/domain/token"
)

// NewBinaryEvaluator is a factory for binary evaluators.
func NewBinaryEvaluator(left interface{}, operator *token.Token, right interface{}) (Evaluator, error) {
	switch operator.Type {
	case token.Minus:
		return &Subtraction{left: left, right: right}, nil
	case token.Plus:
		return &Addition{left: left, right: right}, nil
	case token.Star:
		return &Multiplication{left: left, right: right}, nil
	case token.Slash:
		return &Division{left: left, right: right}, nil
	case token.Modulus:
		return &Modulus{left: left, right: right}, nil
	case token.Greater:
		return &Greater{left: left, right: right}, nil
	case token.GreaterOrEqual:
		return &GreaterOrEqual{left: left, right: right}, nil
	case token.Lower:
		return &Lower{left: left, right: right}, nil
	case token.LowerOrEqual:
		return &LowerOrEqual{left: left, right: right}, nil
	case token.EqualEqual:
		return &Equal{left: left, right: right}, nil
	case token.NotEqual:
		return &Different{left: left, right: right}, nil
	}
	return nil, interr.NewRuntimeError("the operator is not valid", operator.Line)
}

// NewBinaryEvaluator is a factory for unary evaluators.
func NewUnaryEvaluator(operator *token.Token, expression interface{}) (Evaluator, error) {
	switch operator.Type {
	case token.Minus:
		return &MinusNegation{expression: expression}, nil
	case token.Bang:
		return &BangNegation{expression: expression}, nil
	}

	return nil, interr.NewRuntimeError("the operator is not valid", operator.Line)
}
