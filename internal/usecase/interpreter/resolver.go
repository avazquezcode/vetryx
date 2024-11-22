package interpreter

import (
	"fmt"

	"github.com/avazquezcode/govetryx/internal/domain/ast"
	"github.com/avazquezcode/govetryx/internal/domain/token"
	"github.com/avazquezcode/govetryx/internal/domain/types"
)

// Resolver is an important piece of our interpreter, since it resolves the scoping of things.
type Resolver struct {
	interpreter    *Interpreter
	stack          types.Stack
	insideFunction bool // indicates if we are inside a function
	insideLoop     bool // indicates if we are inside a loop
}

func NewResolver(interpreter *Interpreter) *Resolver {
	return &Resolver{
		interpreter: interpreter,
		stack:       types.NewStack(),
	}
}

func (r *Resolver) Resolve(statements []ast.Statement) error {
	for _, statement := range statements {
		err := statement.Accept(r)
		if err != nil {
			return err
		}
	}

	return nil
}

// Simple resolutions

func (r *Resolver) VisitExpressionStatement(v *ast.ExpressionStatement) error {
	_, err := v.Expression.Accept(r)
	return err
}

func (r *Resolver) VisitGroupingExpression(v *ast.GroupingExpression) (interface{}, error) {
	_, err := v.Expression.Accept(r)
	return nil, err
}

func (r *Resolver) VisitLiteralExpression(v *ast.LiteralExpression) (interface{}, error) {
	return nil, nil
}

func (r *Resolver) VisitUnaryExpression(v *ast.UnaryExpression) (interface{}, error) {
	_, err := v.Expression.Accept(r)
	return nil, err
}

func (r *Resolver) VisitPrintStatement(v *ast.PrintStatement) error {
	_, err := v.Expression.Accept(r)
	return err
}

func (r *Resolver) VisitWhileStatement(v *ast.WhileStatement) error {
	_, err := v.Condition.Accept(r)
	if err != nil {
		return err
	}

	r.insideLoop = true

	err = v.Body.Accept(r)

	r.insideLoop = false

	return err
}

func (r *Resolver) VisitBreakStatement(v *ast.BreakStatement) error {
	if !r.insideLoop {
		return fmt.Errorf("cannot execute a break statement outside a loop")
	}
	return nil
}

func (r *Resolver) VisitContinueStatement(v *ast.ContinueStatement) error {
	if !r.insideLoop {
		return fmt.Errorf("cannot execute a continue statement outside a loop")
	}
	return nil
}

func (r *Resolver) VisitBinaryExpression(v *ast.BinaryExpression) (interface{}, error) {
	_, err := v.Left.Accept(r)
	if err != nil {
		return nil, err
	}

	_, err = v.Right.Accept(r)
	return nil, err
}

func (r *Resolver) VisitLogicalExpression(v *ast.LogicalExpression) (interface{}, error) {
	_, err := v.Left.Accept(r)
	if err != nil {
		return nil, err
	}

	_, err = v.Right.Accept(r)
	return nil, err
}

func (r *Resolver) VisitIfStatement(v *ast.IfStatement) error {
	_, err := v.Condition.Accept(r)
	if err != nil {
		return err
	}

	err = v.ThenBlock.Accept(r)
	if err != nil {
		return err
	}

	if v.ElseBlock != nil {
		return v.ElseBlock.Accept(r)
	}

	return nil
}

// Block resolution (start a new scope for the block, and then close it)

func (r *Resolver) VisitBlockStatement(v *ast.BlockStatement) error {
	r.beginScope()
	err := r.Resolve(v.Statements)
	r.endScope()
	return err
}

// Variables resolution

func (r *Resolver) VisitVariableStatement(statement *ast.VariableStatement) error {
	err := r.declare(statement.Name)
	if err != nil {
		return err
	}

	if statement.Value != nil {
		_, err := statement.Value.Accept(r)
		if err != nil {
			return err
		}
	}

	r.define(statement.Name.Lexeme)
	return nil
}

func (r *Resolver) VisitVariableExpression(expression *ast.VariableExpression) (interface{}, error) {
	if r.stack.Length() > 0 {
		initialized, ok := r.stack.Peek().(types.HashMap)[expression.Name.Lexeme]
		if ok && !initialized.(bool) {
			return nil, fmt.Errorf("failed when reading a local variable %q in its initializer", expression.Name.Lexeme)
		}
	}

	return nil, r.resolveLocal(expression, expression.Name.Lexeme)
}

func (r *Resolver) VisitAssignmentExpression(expression *ast.AssignmentExpression) (interface{}, error) {
	_, err := expression.Value.Accept(r)
	if err != nil {
		return nil, err
	}

	return nil, r.resolveLocal(expression, expression.Name.Lexeme)
}

func (r *Resolver) resolveLocal(expression ast.Expression, key string) error {
	lastElementIndex := r.stack.Length() - 1
	for i := lastElementIndex; i >= 0; i-- {
		if r.stack[i].(types.HashMap).Exists(key) {
			// resolve with the key that is closest to the peek of the stack
			r.interpreter.Resolve(expression, lastElementIndex-i)
			return nil
		}
	}

	return nil
}

// Functions resolution

func (r *Resolver) VisitFunctionStatement(statement *ast.FunctionStatement) error {
	err := r.declare(statement.Name)
	if err != nil {
		return err
	}
	r.define(statement.Name.Lexeme)

	return r.resolveFunction(statement)
}

func (r *Resolver) VisitReturnStatement(statement *ast.ReturnStatement) error {
	if !r.insideFunction {
		return fmt.Errorf("cannot return from outside a valid function")
	}

	if statement.Value != nil {
		_, err := statement.Value.Accept(r)
		return err
	}

	return nil
}

func (r *Resolver) VisitCallExpression(expression *ast.CallExpression) (interface{}, error) {
	_, err := expression.Callee.Accept(r)
	if err != nil {
		return nil, err
	}

	for _, arg := range expression.Arguments {
		_, err := arg.Accept(r)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (r *Resolver) resolveFunction(statement *ast.FunctionStatement) error {
	wasInsideFunction := r.insideFunction
	r.beginScope()
	r.insideFunction = true

	for _, param := range statement.Paremeters {
		err := r.declare(param)
		if err != nil {
			return err
		}
		r.define(param.Lexeme)
	}

	err := r.Resolve(statement.Body)

	r.endScope()
	r.insideFunction = wasInsideFunction

	return err
}

// Scope management

func (r *Resolver) beginScope() {
	// Add new scope to the stack.
	r.stack.Push(types.HashMap{})
}

func (r *Resolver) endScope() {
	// Remove last scope from the stack.
	r.stack.Pop()
}

// Declaration and definition

func (r *Resolver) declare(name *token.Token) error {
	if r.stack.Length() == 0 {
		return nil
	}

	currentScope := r.stack.Peek().(types.HashMap)
	if _, exists := currentScope[name.Lexeme]; exists {
		return fmt.Errorf("the variable %q already exists in the scope", name.Lexeme)
	}

	currentScope.Set(name.Lexeme, false)
	return nil
}

func (r *Resolver) define(key string) {
	if r.stack.Length() > 0 {
		currentScope := r.stack.Peek().(types.HashMap)
		currentScope.Set(key, true)
	}
}
