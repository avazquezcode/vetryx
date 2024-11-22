package ast

import "github.com/avazquezcode/govetryx/internal/domain/token"

type (
	// AssignmentExpression is the struct used for assignments.
	AssignmentExpression struct {
		Name  *token.Token
		Value Expression
	}

	// BinaryExpression is the struct used for binary expressions.
	BinaryExpression struct {
		Left     Expression
		Operator *token.Token
		Right    Expression
	}

	// CallExpression is the struct used for calls (eg: a function call).
	CallExpression struct {
		Line      int
		Callee    Expression
		Arguments []Expression
	}

	// GroupingExpression is the struct used for grouping (eg: wrapping an expression with parentheses to indicate a group).
	GroupingExpression struct {
		Expression Expression
	}

	// LiteralExpression is the struct used for literals.
	LiteralExpression struct {
		Value interface{}
	}

	// UnaryExpression is the struct used for unary expressions.
	UnaryExpression struct {
		Operator   *token.Token
		Expression Expression
	}

	// LogicalExpression is the struct used for logical expressions (eg: if condition).
	LogicalExpression struct {
		Left     Expression
		Operator *token.Token
		Right    Expression
	}

	// VariableExpression is the struct used for variable expressions.
	VariableExpression struct {
		Name *token.Token
	}
)

func NewAssignmentExpression(name *token.Token, val Expression) *AssignmentExpression {
	return &AssignmentExpression{
		Name:  name,
		Value: val,
	}
}

func (e *AssignmentExpression) Accept(visitor ExpressionVisitor) (interface{}, error) {
	return visitor.VisitAssignmentExpression(e)
}

func NewBinaryExpression(left Expression, operator *token.Token, right Expression) *BinaryExpression {
	return &BinaryExpression{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}

func (e *BinaryExpression) Accept(visitor ExpressionVisitor) (interface{}, error) {
	return visitor.VisitBinaryExpression(e)
}

func NewCallExpression(line int, callee Expression, args []Expression) *CallExpression {
	return &CallExpression{
		Line:      line,
		Callee:    callee,
		Arguments: args,
	}
}

func (e *CallExpression) Accept(visitor ExpressionVisitor) (interface{}, error) {
	return visitor.VisitCallExpression(e)
}

func NewGroupingExpression(expression Expression) *GroupingExpression {
	return &GroupingExpression{
		Expression: expression,
	}
}

func (e *GroupingExpression) Accept(visitor ExpressionVisitor) (interface{}, error) {
	return visitor.VisitGroupingExpression(e)
}

func NewLiteralExpression(value interface{}) *LiteralExpression {
	return &LiteralExpression{
		Value: value,
	}
}

func (e *LiteralExpression) Accept(visitor ExpressionVisitor) (interface{}, error) {
	return visitor.VisitLiteralExpression(e)
}

func NewLogicalExpression(left Expression, operator *token.Token, right Expression) *LogicalExpression {
	return &LogicalExpression{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}

func (e *LogicalExpression) Accept(visitor ExpressionVisitor) (interface{}, error) {
	return visitor.VisitLogicalExpression(e)
}

func NewUnaryExpression(operator *token.Token, expression Expression) *UnaryExpression {
	return &UnaryExpression{
		Operator:   operator,
		Expression: expression,
	}
}

func (e *UnaryExpression) Accept(visitor ExpressionVisitor) (interface{}, error) {
	return visitor.VisitUnaryExpression(e)
}

func NewVariableExpression(name *token.Token) *VariableExpression {
	return &VariableExpression{
		Name: name,
	}
}

func (e *VariableExpression) Accept(visitor ExpressionVisitor) (interface{}, error) {
	return visitor.VisitVariableExpression(e)
}
