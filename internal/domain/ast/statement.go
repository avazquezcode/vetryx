package ast

import "github.com/avazquezcode/govetryx/internal/domain/token"

type (
	// BlockStatement is the struct used to represent block statements (eg: the block of a while loop).
	BlockStatement struct {
		Statements []Statement
	}

	// ExpressionStatement is the struct used to represent an expression statement.
	ExpressionStatement struct {
		Expression Expression
	}

	// FunctionStatement is the struct used to represent a function statement.
	FunctionStatement struct {
		Name       *token.Token
		Paremeters []*token.Token
		Body       []Statement
	}

	// PrintStatement is the struct used to represent the print statement.
	PrintStatement struct {
		Expression Expression
	}

	// ReturnStatement is the struct used to represent the return statement.
	ReturnStatement struct {
		Line  int
		Value Expression
	}

	// WhileStatement is the struct used to represent the while statement.
	WhileStatement struct {
		Condition Expression
		Body      Statement
	}

	// IfStatement is the struct used to represent an if condition statement.
	IfStatement struct {
		Condition Expression
		ThenBlock Statement
		ElseBlock Statement
	}

	// VariableStatement is the struct used to represent a variable statement.
	VariableStatement struct {
		Name  *token.Token
		Value Expression
	}

	// BreakStatement is the struct used to represent the break statement.
	BreakStatement struct {
		Line int
	}

	// ContinueStatement is the struct used to represent the continue statement.
	ContinueStatement struct {
		Line int
	}
)

func NewBlockStatement(smts []Statement) *BlockStatement {
	return &BlockStatement{
		Statements: smts,
	}
}

func (s *BlockStatement) Accept(visitor StatementVisitor) error {
	return visitor.VisitBlockStatement(s)
}

func NewExpressionStatement(expression Expression) *ExpressionStatement {
	return &ExpressionStatement{
		Expression: expression,
	}
}

func (s *ExpressionStatement) Accept(visitor StatementVisitor) error {
	return visitor.VisitExpressionStatement(s)
}

func NewFunctionStatement(name *token.Token, params []*token.Token, body []Statement) *FunctionStatement {
	return &FunctionStatement{
		Name:       name,
		Paremeters: params,
		Body:       body,
	}
}

func (s *FunctionStatement) Accept(visitor StatementVisitor) error {
	return visitor.VisitFunctionStatement(s)
}

func NewIfStatement(condition Expression, thenBranch Statement, elseBranch Statement) *IfStatement {
	return &IfStatement{
		Condition: condition,
		ThenBlock: thenBranch,
		ElseBlock: elseBranch,
	}
}

func (s *IfStatement) Accept(visitor StatementVisitor) error {
	return visitor.VisitIfStatement(s)
}

func NewPrintStatement(expression Expression) *PrintStatement {
	return &PrintStatement{
		Expression: expression,
	}
}

func (s *PrintStatement) Accept(visitor StatementVisitor) error {
	return visitor.VisitPrintStatement(s)
}

func NewReturnStatement(line int, value Expression) *ReturnStatement {
	return &ReturnStatement{
		Line:  line,
		Value: value,
	}
}

func (s *ReturnStatement) Accept(visitor StatementVisitor) error {
	return visitor.VisitReturnStatement(s)
}

func NewVariableStatement(name *token.Token, initializer Expression) *VariableStatement {
	return &VariableStatement{
		Name:  name,
		Value: initializer,
	}
}

func (s *VariableStatement) Accept(visitor StatementVisitor) error {
	return visitor.VisitVariableStatement(s)
}

func NewWhileStatement(condition Expression, body Statement) *WhileStatement {
	return &WhileStatement{
		Condition: condition,
		Body:      body,
	}
}

func (s *WhileStatement) Accept(visitor StatementVisitor) error {
	return visitor.VisitWhileStatement(s)
}

func NewBreakStatement(line int) *BreakStatement {
	return &BreakStatement{
		Line: line,
	}
}

func (s *BreakStatement) Accept(visitor StatementVisitor) error {
	return visitor.VisitBreakStatement(s)
}

func NewContinueStatement(line int) *ContinueStatement {
	return &ContinueStatement{
		Line: line,
	}
}

func (s *ContinueStatement) Accept(visitor StatementVisitor) error {
	return visitor.VisitContinueStatement(s)
}
