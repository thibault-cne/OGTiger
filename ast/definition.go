package ast

import (
	"ogtiger/parser"
)

type Definition struct {
	Declarations []Ast
	Expressions  []Ast
	Ctx          parser.IDefinitionContext
}

func (e *Definition) Display() string {
	return " letin"
}

func (l *AstCreatorListener) DefinitionEnter(ctx parser.IDefinitionContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) DefinitionExit(ctx parser.IDefinitionContext) {
	expr := &Definition{
		Ctx: ctx,
	}

	for range ctx.AllDeclaration() {
		expr.Declarations = append(expr.Declarations, l.PopAst())
	}

	for range ctx.AllExpression() {
		expr.Expressions = append(expr.Expressions, l.PopAst())
	}

	l.PushAst(expr)
}
