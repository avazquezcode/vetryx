package interpreter

import (
	"errors"
	"fmt"
	"io"

	"github.com/avazquezcode/govetryx/internal/domain/ast"
	"github.com/avazquezcode/govetryx/internal/domain/corerule"
	interr "github.com/avazquezcode/govetryx/internal/domain/error"
	"github.com/avazquezcode/govetryx/internal/domain/evaluator"
	"github.com/avazquezcode/govetryx/internal/domain/token"
	"github.com/avazquezcode/govetryx/internal/domain/types"
)

// Break and continue are defined like errors, to simplify implementation
type Break struct {
	error
}
type Continue struct {
	error
}

type Interpreter struct {
	env    *Env
	global *Env
	local  types.HashMap
	stdout io.Writer
}

// NewInterpreter is a constructor for an interpreter.
func NewInterpreter(stdout io.Writer) *Interpreter {
	global := NewGlobal()

	// Register the native functions in the global environment
	global.Set("sleep", FnSleep{})
	global.Set("clock", FnClock{})
	global.Set("min", FnMin{})
	global.Set("max", FnMax{})

	return &Interpreter{
		env:    global,
		global: global,
		local:  types.HashMap{},
		stdout: stdout,
	}
}

// Interpret is the main method of the interpreter.
// It interprets the code while traversing the AST.
func (i *Interpreter) Interpret(statements []ast.Statement) error {
	for _, statement := range statements {
		err := statement.Accept(i)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *Interpreter) VisitLiteralExpression(expression *ast.LiteralExpression) (interface{}, error) {
	return expression.Value, nil
}

func (i *Interpreter) VisitGroupingExpression(expression *ast.GroupingExpression) (interface{}, error) {
	return expression.Expression.Accept(i)
}

func (i *Interpreter) VisitUnaryExpression(expression *ast.UnaryExpression) (interface{}, error) {
	right, err := expression.Expression.Accept(i)
	if err != nil {
		return nil, interr.NewRuntimeError(err.Error(), expression.Operator.Line)
	}

	evaluator, err := evaluator.NewUnaryEvaluator(expression.Operator, right)
	if err != nil {
		return nil, interr.NewRuntimeError(err.Error(), expression.Operator.Line)
	}

	evaluation, err := evaluator.Evaluate()
	if err != nil {
		return nil, interr.NewRuntimeError(err.Error(), expression.Operator.Line)
	}

	return evaluation, nil
}

func (i *Interpreter) VisitBinaryExpression(expression *ast.BinaryExpression) (interface{}, error) {
	left, err := expression.Left.Accept(i)
	if err != nil {
		return nil, interr.NewRuntimeError(err.Error(), expression.Operator.Line)
	}

	right, err := expression.Right.Accept(i)
	if err != nil {
		return nil, interr.NewRuntimeError(err.Error(), expression.Operator.Line)
	}

	evaluator, err := evaluator.NewBinaryEvaluator(left, expression.Operator, right)
	if err != nil {
		return nil, interr.NewRuntimeError(err.Error(), expression.Operator.Line)
	}

	evaluation, err := evaluator.Evaluate()
	if err != nil {
		return nil, interr.NewRuntimeError(err.Error(), expression.Operator.Line)
	}

	return evaluation, nil
}

func (i *Interpreter) VisitExpressionStatement(statement *ast.ExpressionStatement) error {
	_, err := statement.Expression.Accept(i)
	return err
}

func (i *Interpreter) VisitPrintStatement(statement *ast.PrintStatement) error {
	value, err := statement.Expression.Accept(i)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(i.stdout, corerule.PrintableValue(value))
	if err != nil {
		return fmt.Errorf("failed when printing a value, with err: %w", err)
	}

	return nil
}

func (i *Interpreter) VisitVariableStatement(statement *ast.VariableStatement) error {
	var value interface{}
	var err error

	if statement.Value != nil {
		value, err = statement.Value.Accept(i)
		if err != nil {
			return interr.NewRuntimeError(err.Error(), statement.Name.Line)
		}
	}

	// Set the variable in the environment
	i.env.Set(statement.Name.Lexeme, value)

	return nil
}

func (i *Interpreter) VisitVariableExpression(expression *ast.VariableExpression) (interface{}, error) {
	if i.local.Exists(expression) {
		return i.env.GetAt(i.local.Get(expression).(int), expression.Name.Lexeme)
	}

	return i.global.Get(expression.Name.Lexeme)
}

func (i *Interpreter) VisitAssignmentExpression(expression *ast.AssignmentExpression) (interface{}, error) {
	value, err := expression.Value.Accept(i)
	if err != nil {
		return nil, interr.NewRuntimeError(err.Error(), expression.Name.Line)
	}

	if i.local.Exists(expression) {
		// Means we found it in the local.
		err = i.env.AssignAt(i.local.Get(expression).(int), expression.Name.Lexeme, value)
		if err != nil {
			return nil, interr.NewRuntimeError(err.Error(), expression.Name.Line)
		}
		return value, nil
	}

	// Not in local, so should be in global.
	err = i.global.Assign(expression.Name.Lexeme, value)
	if err != nil {
		return nil, interr.NewRuntimeError(err.Error(), expression.Name.Line)
	}

	return value, nil
}

func (i *Interpreter) VisitBlockStatement(statement *ast.BlockStatement) error {
	return i.executeBlock(statement.Statements, NewLocal(i.env))
}

func (i *Interpreter) executeBlock(statements []ast.Statement, blockEnv *Env) error {
	previousEnv := i.env

	// switch to block env
	i.env = blockEnv

	for _, statement := range statements {
		err := statement.Accept(i)
		if err != nil {
			return err
		}
	}

	// switch back to previous env
	i.env = previousEnv

	return nil
}

func (i *Interpreter) VisitIfStatement(statement *ast.IfStatement) error {
	condition, err := statement.Condition.Accept(i)
	if err != nil {
		return err
	}

	if corerule.IsTrue(condition) {
		return statement.ThenBlock.Accept(i)
	}

	if statement.ElseBlock != nil {
		return statement.ElseBlock.Accept(i)
	}

	return nil
}

func (i *Interpreter) VisitLogicalExpression(expression *ast.LogicalExpression) (interface{}, error) {
	left, err := expression.Left.Accept(i)
	if err != nil {
		return nil, interr.NewRuntimeError(err.Error(), expression.Operator.Line)
	}

	// Implementation of short circuit
	if expression.Operator.Type == token.Or && corerule.IsTrue(left) {
		return left, nil
	}

	if expression.Operator.Type == token.And && !corerule.IsTrue(left) {
		return left, nil
	}

	right, err := expression.Right.Accept(i)
	if err != nil {
		return nil, interr.NewRuntimeError(err.Error(), expression.Operator.Line)
	}

	return right, nil
}

func (i *Interpreter) VisitWhileStatement(statement *ast.WhileStatement) error {
	env := i.env
	for {
		evalCondition, err := statement.Condition.Accept(i)
		if err != nil {
			return err
		}

		if !corerule.IsTrue(evalCondition) {
			break
		}

		err = statement.Body.Accept(i)
		if err != nil {
			i.env = env // This is necessary, so the environment is properly reseted. Same as we do in "executeBlock()".

			if errors.Is(err, Break{}) {
				break
			}

			if errors.Is(err, Continue{}) {
				continue
			}

			// If is not break or continue, we return the real error
			return err
		}
	}

	return nil
}

func (i *Interpreter) VisitCallExpression(expression *ast.CallExpression) (interface{}, error) {
	callee, err := expression.Callee.Accept(i)
	if err != nil {
		return nil, interr.NewRuntimeError(err.Error(), expression.Line)
	}

	// evaluate the arguments
	var arguments []interface{}
	for _, argument := range expression.Arguments {
		evaluatedArgument, err := argument.Accept(i)
		if err != nil {
			return nil, interr.NewRuntimeError(err.Error(), expression.Line)
		}

		arguments = append(arguments, evaluatedArgument)
	}

	function, ok := callee.(callable)
	if !ok {
		return nil, interr.NewRuntimeError("tried to call a non-function", expression.Line)
	}

	if function.Arity() != len(arguments) {
		return nil, interr.NewRuntimeError("the quantity of arguments for the call doesn't match quantity of parameters expected by the function", expression.Line)
	}

	result, err := function.Call(i, arguments)
	if err != nil {
		return nil, interr.NewRuntimeError(err.Error(), expression.Line)
	}

	return result, nil
}

func (i *Interpreter) VisitFunctionStatement(statement *ast.FunctionStatement) error {
	i.env.Set(statement.Name.Lexeme, NewFunction(statement, i.env))
	return nil
}

func (i *Interpreter) VisitReturnStatement(statement *ast.ReturnStatement) error {
	if statement.Value != nil {
		value, err := statement.Value.Accept(i)
		if err != nil {
			return interr.NewRuntimeError(err.Error(), statement.Line)
		}
		// return with value
		panic(NewReturnObj(value))
	}

	// return with no value
	panic(NewReturnObj(nil))
}

func (i *Interpreter) VisitBreakStatement(statement *ast.BreakStatement) error {
	return Break{}
}

func (i *Interpreter) VisitContinueStatement(statement *ast.ContinueStatement) error {
	return Continue{}
}

func (i *Interpreter) Resolve(expression ast.Expression, depth int) {
	i.local.Set(expression, depth)
}
