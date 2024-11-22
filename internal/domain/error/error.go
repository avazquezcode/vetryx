package error

import (
	"fmt"
)

type RuntimeError struct {
	Message string
	Line    int
}

func NewRuntimeError(message string, line int) RuntimeError {
	return RuntimeError{
		Message: message,
		Line:    line,
	}
}

func (r RuntimeError) Error() string {
	if r.Line != 0 {
		return fmt.Sprintf("runtime error occurred at line %d: %s", r.Line, r.Message)
	}

	return r.Message
}
