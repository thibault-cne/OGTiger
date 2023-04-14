package ast

import (
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type Definition struct {
	Declarations []Ast
	Expressions  []Ast
	Ctx          parser.IDefinitionContext
	Type         ttype.TigerType
}

func (e *Definition) ReturnType() ttype.TigerType {
	return e.Type
}

func (e *Definition) Display() string {
	return " letin"
}

func (e *Definition) Draw(prefix string, g *cgraph.Graph) {
	// TODO: Draw the AST
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
