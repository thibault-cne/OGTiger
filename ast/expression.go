package ast

import (
	"ogtiger/parser"
)

type Expression struct {
	Left  Ast
	Right Ast
}

func (e *Expression) Display() string {
	return " expression"
}

func (l *AstCreatorListener) ExprEnter(ctx parser.ExpressionContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) ExprExit(ctx parser.ExpressionContext) {
	// Get back the last element of the stack
	expr := &Expression{}

	if ctx.GetChildCount() == 1 {
		return
	}

	// Get the first term
	left := l.AstStack[len(l.AstStack)-1]       // Take it from the top
	l.AstStack = l.AstStack[:len(l.AstStack)-1] // Remove it
	expr.Left = left                            // Store it

	// Get the second term
	right := l.AstStack[len(l.AstStack)-1]
	l.AstStack = l.AstStack[:len(l.AstStack)-1]
	expr.Right = right

	// Push the new element on the stack
	l.AstStack = append(l.AstStack, expr)
}