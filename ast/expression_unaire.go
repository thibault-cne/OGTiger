package ast

import (
	"ogtiger/parser"

	"github.com/goccy/go-graphviz/cgraph"
)

type ExpressionUnaire struct {
	Expr Ast
	Ctx  parser.IExpressionUnaireContext
}

func (e *ExpressionUnaire) Display() string {
	return " expressionUnaire"
}

func (e *ExpressionUnaire) Draw(g *cgraph.Graph) *cgraph.Node {
	node, _ := g.CreateNode("ExpressionUnaire")

	expr := e.Expr.Draw(g)
	g.CreateEdge("Expr", node, expr)

	return node
}

func (l *AstCreatorListener) ExpressionUnaireEnter(ctx parser.IExpressionUnaireContext) {
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func (l *AstCreatorListener) ExpressionUnaireExit(ctx parser.IExpressionUnaireContext) {
	expressionUnaire := &ExpressionUnaire{
		Ctx: ctx,
	}

	if ctx.GetChildCount() == 1 {
		return
	}

	// Get the first expr
	// Pop the last element of the stack
	expr := l.PopAst()

	expressionUnaire.Expr = expr

	// Push the new element on the stack
	l.PushAst(expressionUnaire)
}
