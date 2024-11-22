// Evaluator is the package that implement the evaluators of the interpreter.
package evaluator

type (
	// Evaluator is the interface implemented by all the evaluators.
	Evaluator interface {
		Evaluate() (interface{}, error)
	}
)
