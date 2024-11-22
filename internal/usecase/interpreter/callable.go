package interpreter

import (
	"github.com/avazquezcode/govetryx/internal/domain/ast"
)

type (
	callable interface {
		Call(interpreter *Interpreter, arguments []interface{}) (interface{}, error)
		Arity() int
	}

	Function struct {
		Declaration *ast.FunctionStatement
		Closure     *Env
	}

	ReturnObj struct {
		Value interface{}
	}
)

func NewReturnObj(value interface{}) *ReturnObj {
	return &ReturnObj{
		Value: value,
	}
}

func NewFunction(declaration *ast.FunctionStatement, closure *Env) *Function {
	return &Function{
		Declaration: declaration,
		Closure:     closure,
	}
}

// Call executes a function call
func (f *Function) Call(interpreter *Interpreter, arguments []interface{}) (result interface{}, err error) {
	// Handle return (using panics)
	defer func() {
		if err := recover(); err != nil {
			if returnObj, ok := err.(*ReturnObj); ok {
				result = returnObj.Value
				return
			}
			panic(err)
		}
	}()

	// Create a local Env for this call
	env := NewLocal(f.Closure)

	// Define the paremeters expected by the function in the local env
	for i, param := range f.Declaration.Paremeters {
		env.Set(param.Lexeme, arguments[i])
	}

	return nil, interpreter.executeBlock(f.Declaration.Body, env)
}

// Arity returns the quantity of parameters defined in the function signature.
func (f *Function) Arity() int {
	return len(f.Declaration.Paremeters)
}
