package evaluator

import (
	"fmt"
	"math"

	"github.com/avazquezcode/govetryx/internal/domain/corerule"
)

type (
	// Different evaluates if left is different than right.
	Different struct {
		left  interface{}
		right interface{}
	}

	// Equal evaluates if left is equal than right.
	Equal struct {
		left  interface{}
		right interface{}
	}

	// GreaterOrEqual evaluates if left is greater or equal than right.
	GreaterOrEqual struct {
		left  interface{}
		right interface{}
	}

	// Greater evaluates if left is greater than right.
	Greater struct {
		left  interface{}
		right interface{}
	}

	// LowerOrEqual evaluates if left is lower or equal than right.
	LowerOrEqual struct {
		left  interface{}
		right interface{}
	}

	// Lower evaluates if left is lower than right.
	Lower struct {
		left  interface{}
		right interface{}
	}

	// Subtraction evaluates the subtraction between left and right (eg: left-right).
	Subtraction struct {
		left  interface{}
		right interface{}
	}

	// Addition evaluates the sum between left and right (eg: left+right).
	// If left and right are strings, it concatenates the values instead.
	Addition struct {
		left  interface{}
		right interface{}
	}

	// Division evaluates the division between left and right (eg: left / right).
	// Division per zero, throws an error.
	Division struct {
		left  interface{}
		right interface{}
	}

	// Multiplication evaluates the multiplication of left and right (eg: left * right).
	Multiplication struct {
		left  interface{}
		right interface{}
	}

	// Modulus evaluates the modulus operation between left and right (eg: left % right).
	Modulus struct {
		left  interface{}
		right interface{}
	}
)

func (a *Different) Evaluate() (interface{}, error) {
	return a.left != a.right, nil
}

func (a *Equal) Evaluate() (interface{}, error) {
	return corerule.IsEqual(a.left, a.right), nil
}

func (a *GreaterOrEqual) Evaluate() (interface{}, error) {
	if isNumber(a.left, a.right) {
		return a.left.(float64) >= a.right.(float64), nil
	}
	return nil, fmt.Errorf("type is invalid")
}

func (a *Greater) Evaluate() (interface{}, error) {
	if isNumber(a.left, a.right) {
		return a.left.(float64) > a.right.(float64), nil
	}
	return nil, fmt.Errorf("type is invalid")
}

func (a *LowerOrEqual) Evaluate() (interface{}, error) {
	if isNumber(a.left, a.right) {
		return a.left.(float64) <= a.right.(float64), nil
	}
	return nil, fmt.Errorf("type is invalid")
}

func (a *Lower) Evaluate() (interface{}, error) {
	if isNumber(a.left, a.right) {
		return a.left.(float64) < a.right.(float64), nil
	}
	return nil, fmt.Errorf("type is invalid")
}

func (a *Subtraction) Evaluate() (interface{}, error) {
	if isNumber := isNumber(a.left, a.right); isNumber {
		return a.left.(float64) - a.right.(float64), nil
	}

	return nil, fmt.Errorf("type is invalid")
}

func (a *Addition) Evaluate() (interface{}, error) {
	if isNumber := isNumber(a.left, a.right); isNumber {
		return a.performSum()
	}

	if isString := isString(a.left, a.right); isString {
		return a.performConcat()
	}

	return nil, fmt.Errorf("type is invalid")

}

func (a *Addition) performSum() (interface{}, error) {
	return a.left.(float64) + a.right.(float64), nil
}

func (a *Addition) performConcat() (interface{}, error) {
	return a.left.(string) + a.right.(string), nil
}

func (a *Division) Evaluate() (interface{}, error) {
	if isNumber(a.left, a.right) {
		if a.right == float64(0) {
			return nil, fmt.Errorf("division per zero")
		}
		return a.left.(float64) / a.right.(float64), nil
	}
	return nil, fmt.Errorf("type is invalid")
}

func (a *Multiplication) Evaluate() (interface{}, error) {
	if isNumber := isNumber(a.left, a.right); isNumber {
		return a.left.(float64) * a.right.(float64), nil
	}

	return nil, fmt.Errorf("type is invalid")
}

func (a *Modulus) Evaluate() (interface{}, error) {
	if isNumber(a.left, a.right) {
		if a.right == float64(0) {
			return nil, fmt.Errorf("division per zero")
		}
		return math.Mod(a.left.(float64), a.right.(float64)), nil
	}
	return nil, fmt.Errorf("type is invalid")
}
