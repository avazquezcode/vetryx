package interpreter

import (
	"fmt"
	"time"
)

type (
	FnClock struct{}
	FnSleep struct{}
	FnMin   struct{}
	FnMax   struct{}
)

func (n FnClock) Arity() int {
	return 0
}

func (n FnClock) Call(interpreter *Interpreter, arguments []interface{}) (interface{}, error) {
	return float64(time.Now().UnixNano()), nil
}

func (n FnSleep) Arity() int {
	return 1
}

func (n FnSleep) Call(interpreter *Interpreter, arguments []interface{}) (interface{}, error) {
	milliSeconds, validFloat := arguments[0].(float64)
	if !validFloat {
		return nil, fmt.Errorf("argument must be a valid float")
	}

	// Sleep
	time.Sleep(time.Duration(float64(time.Millisecond) * milliSeconds))
	return nil, nil
}

func (n FnMin) Arity() int {
	return 2
}

func (n FnMin) Call(interpreter *Interpreter, arguments []interface{}) (interface{}, error) {
	v1, validFloat := arguments[0].(float64)
	if !validFloat {
		return nil, fmt.Errorf("argument must be a valid float")
	}

	v2, validFloat := arguments[1].(float64)
	if !validFloat {
		return nil, fmt.Errorf("argument must be a valid float")
	}

	return min(v1, v2), nil
}

func (n FnMax) Arity() int {
	return 2
}

func (n FnMax) Call(interpreter *Interpreter, arguments []interface{}) (interface{}, error) {
	v1, validFloat := arguments[0].(float64)
	if !validFloat {
		return nil, fmt.Errorf("argument must be a valid float")
	}

	v2, validFloat := arguments[1].(float64)
	if !validFloat {
		return nil, fmt.Errorf("argument must be a valid float")
	}

	return max(v1, v2), nil
}
