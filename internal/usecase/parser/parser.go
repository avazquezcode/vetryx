// This package contains the parser of our interpreter.
// It parses a list of tokens into our AST (Abstract Syntax Tree).
package parser

import (
	"errors"
	"fmt"

	"github.com/avazquezcode/govetryx/internal/domain/ast"
	"github.com/avazquezcode/govetryx/internal/domain/token"
)

// Parser parses a list of scanned tokens into an AST.
type Parser struct {
	tokens  []*token.Token
	current int
}

// NewParser is a constructor for our Parser.
func NewParser(tokens []*token.Token) *Parser {
	return &Parser{tokens: tokens}
}

// Parse is the main method of the parser.
// It parses the list of tokens into our AST, returning a list of statements.
func (p *Parser) Parse() ([]ast.Statement, error) {
	var errors []error
	var statements []ast.Statement

	for !p.isEnd() {
		// Declaration is the entry point of our grammar.
		// A program, is composed by one or many declarations.
		statement, err := p.declaration()
		if err != nil {
			errors = append(errors, err)
			p.synchronize()
			continue
		}
		statements = append(statements, statement)
	}

	if errors != nil {
		return nil, NewParsingErr(errors)
	}

	return statements, nil
}

// declaration is the top of our grammar (program is a set of declarations).
func (p *Parser) declaration() (ast.Statement, error) {
	switch p.peek().Type {
	case token.Fn:
		p.increment()
		return p.function()
	case token.VarDeclarator:
		p.increment()
		return p.variable()
	}
	return p.statement()
}

// function parses a function.
func (p *Parser) function() (ast.Statement, error) {
	functionName, err := p.consume(token.Identifier)
	if err != nil {
		return nil, fmt.Errorf("expected a valid function name: %w", err)
	}

	_, err = p.consume(token.LeftParentheses)
	if err != nil {
		return nil, fmt.Errorf("expected '(' after function name: %w", err)
	}

	var parameters []*token.Token
	if !p.is(token.RightParentheses) {
		parameter, err := p.consume(token.Identifier)
		if err != nil {
			return nil, fmt.Errorf("expected a valid parameter after '(': %w", err)
		}
		parameters = append(parameters, parameter)

		// If we have a comma, it means we should expect more than one parameter.
		for p.is(token.Comma) {
			p.increment() // skip the comma

			param, err := p.consume(token.Identifier)
			if err != nil {
				return nil, fmt.Errorf("expected a valid parameter after ',': %w", err)
			}
			parameters = append(parameters, param)
		}
	}

	_, err = p.consume(token.RightParentheses)
	if err != nil {
		return nil, fmt.Errorf("expected ')' after the function parameters list: %w", err)
	}

	_, err = p.consume(token.LeftBrace)
	if err != nil {
		return nil, fmt.Errorf("expected '{' before the function body: %w", err)
	}

	body, err := p.block()
	if err != nil {
		return nil, err
	}

	return ast.NewFunctionStatement(functionName, parameters, body), nil
}

// block parses a block.
func (p *Parser) block() ([]ast.Statement, error) {
	var statements []ast.Statement

	for !p.is(token.RightBrace) && !p.isEnd() {
		statement, err := p.declaration()
		if err != nil {
			return nil, err
		}
		statements = append(statements, statement)
	}

	_, err := p.consume(token.RightBrace)
	if err != nil {
		return nil, fmt.Errorf("expected a '}' to close the block: %w", err)
	}

	return statements, nil
}

// variable parses a variable declaration.
func (p *Parser) variable() (ast.Statement, error) {
	name, err := p.consume(token.Identifier)
	if err != nil {
		return nil, fmt.Errorf("expected a valid variable name: %w", err)
	}

	var variableValue ast.Expression
	if p.is(token.Equal) {
		p.increment() // skip the equal

		variableValue, err = p.expression()
		if err != nil {
			return nil, err
		}
	}

	_, err = p.consume(token.Semicolon)
	if err != nil {
		return nil, fmt.Errorf("expected a ';' after the variable declaration: %w", err)
	}

	return ast.NewVariableStatement(name, variableValue), nil
}

// statement parses a statement.
func (p *Parser) statement() (ast.Statement, error) {
	if !p.isEnd() && p.peekNext().Type == token.VarShortDeclarator {
		return p.varShortDeclaratorStatement()
	}

	switch p.peek().Type {
	case token.If:
		p.increment()
		return p.ifStatement()
	case token.While:
		p.increment()
		return p.whileStatement()
	case token.Return:
		p.increment()
		return p.returnStatement()
	case token.Print:
		p.increment()
		return p.printStatement()
	case token.Break:
		p.increment()
		return p.breakStatement()
	case token.Continue:
		p.increment()
		return p.continueStatement()
	case token.LeftBrace:
		p.increment()

		statement, err := p.block()
		if err != nil {
			return nil, err
		}
		return ast.NewBlockStatement(statement), nil
	}

	return p.expressionStatement()
}

// ifStmt parses an if statement.
func (p *Parser) ifStatement() (ast.Statement, error) {
	condition, err := p.expression()
	if err != nil {
		return nil, err
	}

	thenBlock, err := p.statement()
	if err != nil {
		return nil, err
	}

	var elseBlock ast.Statement
	if p.is(token.Else) {
		p.increment() // skip the else word

		elseBlock, err = p.statement()
		if err != nil {
			return nil, err
		}
	}

	return ast.NewIfStatement(condition, thenBlock, elseBlock), nil
}

// whileStmt parses a while statement.
func (p *Parser) whileStatement() (ast.Statement, error) {
	condition, err := p.expression()
	if err != nil {
		return nil, err
	}

	body, err := p.statement()
	if err != nil {
		return nil, err
	}

	return ast.NewWhileStatement(condition, body), nil
}

// varShortDeclaratorStmt parses a short declaration of a variable.
func (p *Parser) varShortDeclaratorStatement() (ast.Statement, error) {
	name, err := p.consume(token.Identifier)
	if err != nil {
		return nil, fmt.Errorf("expected a valid variable name: %w", err)
	}

	_, err = p.consume(token.VarShortDeclarator)
	if err != nil {
		return nil, fmt.Errorf("expected a valid short declarator assignment operator: %w", err)
	}

	initializer, err := p.expression()
	if err != nil {
		return nil, err
	}

	_, err = p.consume(token.Semicolon)
	if err != nil {
		return nil, fmt.Errorf("expected a ';' after the variable declaration: %w", err)
	}

	return ast.NewVariableStatement(name, initializer), nil
}

// returnStmt parses a return statement.
func (p *Parser) returnStatement() (ast.Statement, error) {
	returnLine := p.previous().Line
	var value ast.Expression
	var err error

	if !p.is(token.Semicolon) {
		value, err = p.expression()
		if err != nil {
			return nil, err
		}
	}

	_, err = p.consume(token.Semicolon)
	if err != nil {
		return nil, fmt.Errorf("expected a ';' after the return: %w", err)
	}

	return ast.NewReturnStatement(returnLine, value), nil
}

// breakStatement parses a break statement.
func (p *Parser) breakStatement() (ast.Statement, error) {
	breakLine := p.previous().Line
	_, err := p.consume(token.Semicolon)
	if err != nil {
		return nil, fmt.Errorf("expected a ';' after the break: %w", err)
	}

	return ast.NewBreakStatement(breakLine), nil
}

// continueStatement parses a continue statement.
func (p *Parser) continueStatement() (ast.Statement, error) {
	continueLine := p.previous().Line
	_, err := p.consume(token.Semicolon)
	if err != nil {
		return nil, fmt.Errorf("expected a ';' after the continue: %w", err)
	}

	return ast.NewContinueStatement(continueLine), nil
}

// print parses a print statement.
func (p *Parser) printStatement() (ast.Statement, error) {
	value, err := p.expression()
	if err != nil {
		return nil, err
	}

	_, err = p.consume(token.Semicolon)
	if err != nil {
		return nil, fmt.Errorf("expected a ';' after the print statement: %w", err)
	}

	return ast.NewPrintStatement(value), nil
}

// exprStmt parses an expression statement.
func (p *Parser) expressionStatement() (ast.Statement, error) {
	value, err := p.expression()
	if err != nil {
		return nil, err
	}

	_, err = p.consume(token.Semicolon)
	if err != nil {
		return nil, fmt.Errorf("expected a ';' after the expression: %w", err)
	}

	return ast.NewExpressionStatement(value), nil
}

func (p *Parser) expression() (ast.Expression, error) {
	return p.assignment()
}

func (p *Parser) assignment() (ast.Expression, error) {
	expression, err := p.or()
	if err != nil {
		return nil, err
	}

	if !p.is(token.Equal) && !p.is(token.VarShortDeclarator) {
		return expression, nil
	}

	p.increment()

	value, err := p.assignment()
	if err != nil {
		return nil, fmt.Errorf("failed when parsing the assignment value: %w", err)
	}

	name, ok := expression.(*ast.VariableExpression)
	if !ok {
		return nil, errors.New("invalid assignment")
	}

	return ast.NewAssignmentExpression(name.Name, value), nil
}

func (p *Parser) or() (ast.Expression, error) {
	expression, err := p.and()
	if err != nil {
		return nil, err
	}

	for p.is(token.Or) {
		p.increment() // skip the "||"
		operator := p.previous()

		right, err := p.and()
		if err != nil {
			return nil, err
		}

		expression = ast.NewLogicalExpression(expression, operator, right)
	}

	return expression, nil
}

func (p *Parser) and() (ast.Expression, error) {
	expression, err := p.equality()
	if err != nil {
		return nil, err
	}

	for p.is(token.And) {
		p.increment() // skip the "&&"

		operator := p.previous()

		right, err := p.equality()
		if err != nil {
			return nil, err
		}

		expression = ast.NewLogicalExpression(expression, operator, right)
	}

	return expression, nil
}

// equality parses an equality.
func (p *Parser) equality() (ast.Expression, error) {
	expression, err := p.comparison()
	if err != nil {
		return nil, err
	}

	for p.is(token.NotEqual, token.EqualEqual) {
		p.increment()

		operator := p.previous()

		right, err := p.comparison()
		if err != nil {
			return nil, err
		}

		expression = ast.NewBinaryExpression(expression, operator, right)
	}

	return expression, nil
}

// comparison parses a comparison.
func (p *Parser) comparison() (ast.Expression, error) {
	expression, err := p.term()
	if err != nil {
		return nil, err
	}

	for p.is(token.Lower, token.LowerOrEqual, token.Greater, token.GreaterOrEqual) {
		p.increment() // skip the "< or > or >= or <="

		operator := p.previous()

		right, err := p.term()
		if err != nil {
			return nil, err
		}

		expression = ast.NewBinaryExpression(expression, operator, right)
	}

	return expression, nil
}

// term parses a term.
func (p *Parser) term() (ast.Expression, error) {
	expression, err := p.factor()
	if err != nil {
		return nil, err
	}

	for p.is(token.Minus, token.Plus) {
		p.increment() // skip the "-" or "+"

		operator := p.previous()

		right, err := p.factor()
		if err != nil {
			return nil, err
		}
		expression = ast.NewBinaryExpression(expression, operator, right)
	}

	return expression, nil
}

// factor parses a factor.
func (p *Parser) factor() (ast.Expression, error) {
	expression, err := p.unary()
	if err != nil {
		return nil, err
	}

	for p.is(token.Slash, token.Star, token.Modulus) {
		p.increment() // skip the "/", "*" or "%"

		operator := p.previous()

		right, err := p.unary()
		if err != nil {
			return nil, err
		}

		expression = ast.NewBinaryExpression(expression, operator, right)
	}

	return expression, nil
}

// unary parses a unary.
func (p *Parser) unary() (ast.Expression, error) {
	if p.is(token.Bang, token.Minus) {
		p.increment() // Skip the "!" or "-"

		operator := p.previous()

		expression, err := p.unary()
		if err != nil {
			return nil, err
		}

		return ast.NewUnaryExpression(operator, expression), nil
	}

	return p.call()
}

// call parses a call.
func (p *Parser) call() (ast.Expression, error) {
	expression, err := p.primary()
	if err != nil {
		return nil, err
	}

	for p.is(token.LeftParentheses) {
		p.increment() // skip the parentheses
		expression, err = p.parseCall(expression)
		if err != nil {
			return nil, err
		}
	}

	return expression, nil
}

// primary parses a primary.
func (p *Parser) primary() (ast.Expression, error) {
	if p.is(token.Number, token.String) {
		p.increment()
		return ast.NewLiteralExpression(p.previous().Literal), nil
	}
	if p.is(token.True) {
		p.increment()
		return ast.NewLiteralExpression(true), nil
	}
	if p.is(token.False) {
		p.increment()
		return ast.NewLiteralExpression(false), nil
	}
	if p.is(token.Null) {
		p.increment()
		return ast.NewLiteralExpression(nil), nil
	}
	if p.is(token.Identifier) {
		p.increment()
		return ast.NewVariableExpression(p.previous()), nil
	}

	// Handle grouping
	if p.is(token.LeftParentheses) {
		p.increment()

		expression, err := p.expression()
		if err != nil {
			return nil, err
		}

		_, err = p.consume(token.RightParentheses)
		if err != nil {
			return nil, fmt.Errorf("expected ')' after the expression: %w", err)
		}

		return ast.NewGroupingExpression(expression), nil
	}

	return nil, fmt.Errorf("expected an expression")
}

func (p *Parser) parseCall(callee ast.Expression) (ast.Expression, error) {
	var arguments []ast.Expression

	if !p.is(token.RightParentheses) {
		argument, err := p.expression()
		if err != nil {
			return nil, fmt.Errorf("failed when parsing call argument: %w", err)
		}

		arguments = append(arguments, argument)

		for p.is(token.Comma) {
			p.increment() // skip the comma
			argument, err := p.expression()
			if err != nil {
				return nil, fmt.Errorf("failed when parsing call argument: %w", err)
			}

			arguments = append(arguments, argument)
		}
	}

	closingParen, err := p.consume(token.RightParentheses)
	if err != nil {
		return nil, fmt.Errorf("expected a closing ')' after the call arguments: %w", err)
	}

	return ast.NewCallExpression(closingParen.Line, callee, arguments), nil
}

func (p *Parser) previous() *token.Token {
	return p.tokens[p.current-1]
}

func (p *Parser) is(tokenTypes ...token.Type) bool {
	for _, tokenType := range tokenTypes {
		if p.peek().Type == tokenType {
			return true
		}
	}

	return false
}

func (p *Parser) increment() {
	if p.isEnd() {
		return
	}
	p.current++
}

func (p *Parser) isEnd() bool {
	return p.peek().Type == token.EOF
}

func (p *Parser) peek() *token.Token {
	return p.tokens[p.current]
}

func (p *Parser) peekNext() *token.Token {
	return p.tokens[p.current+1]
}

func (p *Parser) consume(expected token.Type) (*token.Token, error) {
	if p.is(expected) {
		p.increment()
		return p.previous(), nil
	}

	return nil, fmt.Errorf("unexpected token at line %d", p.peek().Line)
}

func (p *Parser) synchronize() {
	p.increment()

	for !p.isEnd() {
		if p.previous().Type == token.Semicolon {
			return
		}

		switch p.peek().Type {
		case
			token.Fn,
			token.VarDeclarator,
			token.If,
			token.While,
			token.Return,
			token.Print:
			return
		}

		p.increment()
	}
}
