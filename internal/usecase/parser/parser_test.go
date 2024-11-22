package parser_test

import (
	"bytes"
	"testing"

	"github.com/avazquezcode/govetryx/internal/domain/ast"
	"github.com/avazquezcode/govetryx/internal/domain/token"
	"github.com/avazquezcode/govetryx/internal/usecase/parser"
	"github.com/avazquezcode/govetryx/internal/usecase/scanner"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := map[string]struct {
		src         string
		expected    []ast.Statement
		expectedErr bool
	}{
		"empty src": {
			src:      "",
			expected: nil,
		},
		"variable declaration without value": {
			src: "dec a;",
			expected: []ast.Statement{
				ast.NewVariableStatement(
					token.NewToken(token.Identifier, "a", nil, 1),
					nil),
			},
		},
		"variable declaration": {
			src: "dec a = 1;",
			expected: []ast.Statement{
				ast.NewVariableStatement(
					token.NewToken(token.Identifier, "a", nil, 1),
					ast.NewLiteralExpression(float64(1))),
			},
		},
		"variable declaration of string": {
			src: "dec a = \"hello\";",
			expected: []ast.Statement{
				ast.NewVariableStatement(
					token.NewToken(token.Identifier, "a", nil, 1),
					ast.NewLiteralExpression("hello")),
			},
		},
		"variable declaration of boolean (true)": {
			src: "dec a = true;",
			expected: []ast.Statement{
				ast.NewVariableStatement(
					token.NewToken(token.Identifier, "a", nil, 1),
					ast.NewLiteralExpression(true)),
			},
		},
		"variable declaration of boolean (false)": {
			src: "dec a = false;",
			expected: []ast.Statement{
				ast.NewVariableStatement(
					token.NewToken(token.Identifier, "a", nil, 1),
					ast.NewLiteralExpression(false)),
			},
		},
		"variable declaration of null": {
			src: "dec a = null;",
			expected: []ast.Statement{
				ast.NewVariableStatement(
					token.NewToken(token.Identifier, "a", nil, 1),
					ast.NewLiteralExpression(nil)),
			},
		},
		"sum of two numbers": {
			src: "1 + 1;",
			expected: []ast.Statement{
				ast.NewExpressionStatement(
					ast.NewBinaryExpression(ast.NewLiteralExpression(float64(1)),
						token.NewToken(token.Plus, "+", nil, 1),
						ast.NewLiteralExpression(float64(1)))),
			},
		},
		"subtraction of two numbers": {
			src: "1 - 1;",
			expected: []ast.Statement{
				ast.NewExpressionStatement(
					ast.NewBinaryExpression(ast.NewLiteralExpression(float64(1)),
						token.NewToken(token.Minus, "-", nil, 1),
						ast.NewLiteralExpression(float64(1)))),
			},
		},
		"multiplication of two numbers": {
			src: "1 * 1;",
			expected: []ast.Statement{
				ast.NewExpressionStatement(
					ast.NewBinaryExpression(ast.NewLiteralExpression(float64(1)),
						token.NewToken(token.Star, "*", nil, 1),
						ast.NewLiteralExpression(float64(1)))),
			},
		},
		"division of two numbers": {
			src: "1 / 1;",
			expected: []ast.Statement{
				ast.NewExpressionStatement(
					ast.NewBinaryExpression(ast.NewLiteralExpression(float64(1)),
						token.NewToken(token.Slash, "/", nil, 1),
						ast.NewLiteralExpression(float64(1)))),
			},
		},
		"modulus operation between two numbers": {
			src: "1 % 1;",
			expected: []ast.Statement{
				ast.NewExpressionStatement(
					ast.NewBinaryExpression(ast.NewLiteralExpression(float64(1)),
						token.NewToken(token.Modulus, "%", nil, 1),
						ast.NewLiteralExpression(float64(1)))),
			},
		},
		"equality": {
			src: "1 == 1;",
			expected: []ast.Statement{
				ast.NewExpressionStatement(
					ast.NewBinaryExpression(ast.NewLiteralExpression(float64(1)),
						token.NewToken(token.EqualEqual, "==", nil, 1),
						ast.NewLiteralExpression(float64(1)))),
			},
		},
		"not equal": {
			src: "1 <> 1;",
			expected: []ast.Statement{
				ast.NewExpressionStatement(
					ast.NewBinaryExpression(ast.NewLiteralExpression(float64(1)),
						token.NewToken(token.NotEqual, "<>", nil, 1),
						ast.NewLiteralExpression(float64(1)))),
			},
		},
		"greater": {
			src: "1 > 1;",
			expected: []ast.Statement{
				ast.NewExpressionStatement(
					ast.NewBinaryExpression(ast.NewLiteralExpression(float64(1)),
						token.NewToken(token.Greater, ">", nil, 1),
						ast.NewLiteralExpression(float64(1)))),
			},
		},
		"greater or equal": {
			src: "1 >= 1;",
			expected: []ast.Statement{
				ast.NewExpressionStatement(
					ast.NewBinaryExpression(ast.NewLiteralExpression(float64(1)),
						token.NewToken(token.GreaterOrEqual, ">=", nil, 1),
						ast.NewLiteralExpression(float64(1)))),
			},
		},
		"lower": {
			src: "1 < 1;",
			expected: []ast.Statement{
				ast.NewExpressionStatement(
					ast.NewBinaryExpression(ast.NewLiteralExpression(float64(1)),
						token.NewToken(token.Lower, "<", nil, 1),
						ast.NewLiteralExpression(float64(1)))),
			},
		},
		"lower or equal": {
			src: "1 <= 1;",
			expected: []ast.Statement{
				ast.NewExpressionStatement(
					ast.NewBinaryExpression(
						ast.NewLiteralExpression(float64(1)),
						token.NewToken(token.LowerOrEqual, "<=", nil, 1),
						ast.NewLiteralExpression(float64(1)))),
			},
		},
		"not": {
			src: "!1;",
			expected: []ast.Statement{
				ast.NewExpressionStatement(
					ast.NewUnaryExpression(
						token.NewToken(token.Bang, "!", nil, 1),
						ast.NewLiteralExpression(float64(1)))),
			},
		},
		"negation": {
			src: "-1;",
			expected: []ast.Statement{
				ast.NewExpressionStatement(
					ast.NewUnaryExpression(
						token.NewToken(token.Minus, "-", nil, 1),
						ast.NewLiteralExpression(float64(1)))),
			},
		},
		"grouping": {
			src: "(1+1)*2;",
			expected: []ast.Statement{
				ast.NewExpressionStatement(
					ast.NewBinaryExpression(
						ast.NewGroupingExpression(ast.NewBinaryExpression(
							ast.NewLiteralExpression(float64(1)),
							token.NewToken(token.Plus, "+", nil, 1),
							ast.NewLiteralExpression(float64(1)))),
						token.NewToken(token.Star, "*", nil, 1),
						ast.NewLiteralExpression(float64(2)))),
			},
		},
		"and": {
			src: "(1 && 2);",
			expected: []ast.Statement{
				ast.NewExpressionStatement(
					ast.NewGroupingExpression(
						ast.NewLogicalExpression(
							ast.NewLiteralExpression(float64(1)),
							token.NewToken(token.And, "&&", nil, 1),
							ast.NewLiteralExpression(float64(2))))),
			},
		},
		"or": {
			src: "(1 || 2);",
			expected: []ast.Statement{
				ast.NewExpressionStatement(
					ast.NewGroupingExpression(
						ast.NewLogicalExpression(
							ast.NewLiteralExpression(float64(1)),
							token.NewToken(token.Or, "||", nil, 1),
							ast.NewLiteralExpression(float64(2))))),
			},
		},
		"assignment": {
			src: "a = 1;",
			expected: []ast.Statement{
				ast.NewExpressionStatement(
					ast.NewAssignmentExpression(
						token.NewToken(token.Identifier, "a", nil, 1),
						ast.NewLiteralExpression(float64(1)))),
			},
		},
		"assignment with variable declaration (short var declarator)": {
			src: "a := \"hello\";",
			expected: []ast.Statement{
				ast.NewVariableStatement(
					token.NewToken(token.Identifier, "a", nil, 1),
					ast.NewLiteralExpression("hello")),
			},
		},
		"if condition": {
			src: "if 1 == 1 {}",
			expected: []ast.Statement{
				ast.NewIfStatement(
					ast.NewBinaryExpression(
						ast.NewLiteralExpression(float64(1)),
						token.NewToken(token.EqualEqual, "==", nil, 1),
						ast.NewLiteralExpression(float64(1))),
					ast.NewBlockStatement(nil),
					nil),
			},
		},
		"if condition (with wrapping parentheses)": {
			src: "if (1 == 1) {}",
			expected: []ast.Statement{
				ast.NewIfStatement(
					ast.NewGroupingExpression(
						ast.NewBinaryExpression(
							ast.NewLiteralExpression(float64(1)),
							token.NewToken(token.EqualEqual, "==", nil, 1),
							ast.NewLiteralExpression(float64(1))),
					),
					ast.NewBlockStatement(nil),
					nil),
			},
		},
		"if else condition": {
			src: "if 1 == 1 {} else {}",
			expected: []ast.Statement{
				ast.NewIfStatement(
					ast.NewBinaryExpression(
						ast.NewLiteralExpression(float64(1)),
						token.NewToken(token.EqualEqual, "==", nil, 1),
						ast.NewLiteralExpression(float64(1))),
					ast.NewBlockStatement(nil),
					ast.NewBlockStatement(nil)),
			},
		},
		"while loop": {
			src: "while 1 == 1 {}",
			expected: []ast.Statement{
				ast.NewWhileStatement(
					ast.NewBinaryExpression(
						ast.NewLiteralExpression(float64(1)),
						token.NewToken(token.EqualEqual, "==", nil, 1),
						ast.NewLiteralExpression(float64(1))),
					ast.NewBlockStatement(nil)),
			},
		},
		"while loop (with paren)": {
			src: "while (1 == 1) {}",
			expected: []ast.Statement{
				ast.NewWhileStatement(
					ast.NewGroupingExpression(
						ast.NewBinaryExpression(
							ast.NewLiteralExpression(float64(1)),
							token.NewToken(token.EqualEqual, "==", nil, 1),
							ast.NewLiteralExpression(float64(1))),
					),
					ast.NewBlockStatement(nil)),
			},
		},
		"while loop with break": {
			src: "while 1 == 1 {break;}",
			expected: []ast.Statement{
				ast.NewWhileStatement(
					ast.NewBinaryExpression(
						ast.NewLiteralExpression(float64(1)),
						token.NewToken(token.EqualEqual, "==", nil, 1),
						ast.NewLiteralExpression(float64(1))),
					ast.NewBlockStatement(
						[]ast.Statement{
							ast.NewBreakStatement(1),
						},
					)),
			},
		},
		"while loop with continue": {
			src: "while 1 == 1 {continue;}",
			expected: []ast.Statement{
				ast.NewWhileStatement(
					ast.NewBinaryExpression(
						ast.NewLiteralExpression(float64(1)),
						token.NewToken(token.EqualEqual, "==", nil, 1),
						ast.NewLiteralExpression(float64(1))),
					ast.NewBlockStatement(
						[]ast.Statement{
							ast.NewContinueStatement(1),
						},
					)),
			},
		},
		"function declaration": {
			src: "fn a() {}",
			expected: []ast.Statement{
				ast.NewFunctionStatement(
					token.NewToken(token.Identifier, "a", nil, 1),
					nil,
					nil),
			},
		},
		"function declaration with parameters": {
			src: "fn a(b, c) {}",
			expected: []ast.Statement{
				ast.NewFunctionStatement(
					token.NewToken(token.Identifier, "a", nil, 1),
					[]*token.Token{
						token.NewToken(token.Identifier, "b", nil, 1),
						token.NewToken(token.Identifier, "c", nil, 1),
					},
					nil),
			},
		},
		"function declaration with some body": {
			src: "fn a(b, c) { dec a = 1; }",
			expected: []ast.Statement{
				ast.NewFunctionStatement(
					token.NewToken(token.Identifier, "a", nil, 1),
					[]*token.Token{
						token.NewToken(token.Identifier, "b", nil, 1),
						token.NewToken(token.Identifier, "c", nil, 1),
					},
					[]ast.Statement{
						ast.NewVariableStatement(
							token.NewToken(token.Identifier, "a", nil, 1),
							ast.NewLiteralExpression(float64(1))),
					}),
			},
		},
		"function declaration with some body and return": {
			src: "fn a(b, c) { dec a = 1; return a; }",
			expected: []ast.Statement{
				ast.NewFunctionStatement(
					token.NewToken(token.Identifier, "a", nil, 1),
					[]*token.Token{
						token.NewToken(token.Identifier, "b", nil, 1),
						token.NewToken(token.Identifier, "c", nil, 1),
					},
					[]ast.Statement{
						ast.NewVariableStatement(
							token.NewToken(token.Identifier, "a", nil, 1),
							ast.NewLiteralExpression(float64(1))),
						ast.NewReturnStatement(
							1,
							ast.NewVariableExpression(token.NewToken(token.Identifier, "a", nil, 1))),
					}),
			},
		},
		"print": {
			src: "print 1;",
			expected: []ast.Statement{
				ast.NewPrintStatement(
					ast.NewLiteralExpression(float64(1))),
			},
		},
		"call function": {
			src: "a();",
			expected: []ast.Statement{
				ast.NewExpressionStatement(
					ast.NewCallExpression(
						1,
						ast.NewVariableExpression(token.NewToken(token.Identifier, "a", nil, 1)),
						nil)),
			},
		},
		"call function with args": {
			src: "a(b, c);",
			expected: []ast.Statement{
				ast.NewExpressionStatement(
					ast.NewCallExpression(
						1,
						ast.NewVariableExpression(token.NewToken(token.Identifier, "a", nil, 1)),
						[]ast.Expression{
							ast.NewVariableExpression(token.NewToken(token.Identifier, "b", nil, 1)),
							ast.NewVariableExpression(token.NewToken(token.Identifier, "c", nil, 1)),
						})),
			},
		},
		// ERRORS SECTION
		"missing identifier after fn declaration": {
			src:         "fn ()",
			expectedErr: true,
		},
		"invalid identifier for function name": {
			src:         "fn 123()",
			expectedErr: true,
		},
		"missing opening parentheses after function name": {
			src:         "fn a",
			expectedErr: true,
		},
		"missing parameter after fn declaration": {
			src:         "fn a(",
			expectedErr: true,
		},
		"missing next parameter after comma": {
			src:         "fn a(b,)",
			expectedErr: true,
		},
		"missing closing parentheses for closing parameters declaration of fn": {
			src:         "fn a(b,c",
			expectedErr: true,
		},
		"missing function body": {
			src:         "fn a(b)",
			expectedErr: true,
		},
		"missing block closing brace after fn declaration": {
			src:         "fn a(b) {",
			expectedErr: true,
		},
		"missing variable name": {
			src:         "dec = 1;",
			expectedErr: true,
		},
		"invalid variable initializer": {
			src:         "dec a = #;",
			expectedErr: true,
		},
		"invalid right operand in comparison": {
			src:         "a > %;",
			expectedErr: true,
		},
		"invalid right operand in equality comparison": {
			src:         "a == %;",
			expectedErr: true,
		},
		"missing opening parentheses after while": {
			src:         "while",
			expectedErr: true,
		},
		"missing closing parentheses after while condition": {
			src:         "while (1==1",
			expectedErr: true,
		},
		"missing while condition": {
			src:         "while ()",
			expectedErr: true,
		},
		"wrong right operand on unary operation": {
			src:         "!>",
			expectedErr: true,
		},
		"wrong right operand on factor operation": {
			src:         "1 * >",
			expectedErr: true,
		},
		"wrong right operand on term operation": {
			src:         "1 + >",
			expectedErr: true,
		},
		"wrong right operand on AND logic operation": {
			src:         "1 && >",
			expectedErr: true,
		},
		"wrong right operand on OR logic operation": {
			src:         "1 || >",
			expectedErr: true,
		},
		"wrong return statement": {
			src:         "fn a() {return}",
			expectedErr: true,
		},
		"wrong first argument in fn call": {
			src:         "a(>);",
			expectedErr: true,
		},
		"missing argument after comma in fn call": {
			src:         "a(1,);",
			expectedErr: true,
		},
		"missing right parentheses in fn call with args": {
			src:         "a(1",
			expectedErr: true,
		},
		"wrong assignment": {
			src:         "a = ;",
			expectedErr: true,
		},
		"wrong assignment 2 (invalid value)": {
			src:         "a = >;",
			expectedErr: true,
		},
		"printing something that doesn't make sense": {
			src:         "print >;",
			expectedErr: true,
		},
		"invalid assignment target": {
			src:         "1 = 1;",
			expectedErr: true,
		},
		"missing block closing brace": {
			src:         "{",
			expectedErr: true,
		},
		"missing closing parentheses": {
			src:         "(",
			expectedErr: true,
		},
	}

	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			lexer := scanner.NewScanner(bytes.Runes(strToBytes(test.src)))
			tokens, _ := lexer.Scan() // We are generating a src that we know is valid, so no need for handling error here
			parser := parser.NewParser(tokens)
			statements, err := parser.Parse()

			if test.expectedErr {
				assert.NotNil(t, err)
				return
			}

			assert.Equal(t, test.expected, statements)
			assert.Nil(t, err)
		})
	}
}

func strToBytes(str string) []byte {
	return []byte(str)
}
