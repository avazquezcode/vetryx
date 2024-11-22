// This package contains the Abstract Syntax Tree (AST) used in the interpreter.
// The AST is implemented using the Visitor pattern.
package ast

// Expression is an interface for an expression.
type Expression interface {
	Accept(visitor ExpressionVisitor) (interface{}, error)
}

// Statement is an interface for a statement.
type Statement interface {
	Accept(visitor StatementVisitor) error
}

// ExpressionVisitor ...
type ExpressionVisitor interface {
	VisitGroupingExpression(expression *GroupingExpression) (interface{}, error)
	VisitUnaryExpression(expression *UnaryExpression) (interface{}, error)
	VisitBinaryExpression(expression *BinaryExpression) (interface{}, error)
	VisitAssignmentExpression(expression *AssignmentExpression) (interface{}, error)
	VisitVariableExpression(expression *VariableExpression) (interface{}, error)
	VisitLogicalExpression(expression *LogicalExpression) (interface{}, error)
	VisitLiteralExpression(expression *LiteralExpression) (interface{}, error)
	VisitCallExpression(expression *CallExpression) (interface{}, error)
}

// StatementVisitor ...
type StatementVisitor interface {
	VisitExpressionStatement(statement *ExpressionStatement) error
	VisitReturnStatement(statement *ReturnStatement) error
	VisitVariableStatement(statement *VariableStatement) error
	VisitFunctionStatement(statement *FunctionStatement) error
	VisitIfStatement(statement *IfStatement) error
	VisitPrintStatement(statement *PrintStatement) error
	VisitBlockStatement(statement *BlockStatement) error
	VisitWhileStatement(statement *WhileStatement) error
	VisitBreakStatement(statement *BreakStatement) error
	VisitContinueStatement(statement *ContinueStatement) error
}
