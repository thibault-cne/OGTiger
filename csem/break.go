package ast

import (
	"fmt"
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type Break struct {
	Ctx  parser.BreakContext
	Type *ttype.TigerType
}

func (e *Break) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *Break) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("break")

	return node
}

func (l *SemControlListener) BreakEnter(ctx parser.BreakContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *SemControlListener) BreakExit(ctx parser.BreakContext) {
	// Get back the last element of the stack
	it := &Break{
		Ctx: ctx,
	}

	l.PushAst(it)
}
