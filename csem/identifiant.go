package ast

import (
	"fmt"
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type Identifiant struct {
	Id   string
	Ctx  parser.IIdentifiantContext
	Type *ttype.TigerType
}

func (e *Identifiant) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *Identifiant) Draw(g *cgraph.Graph) *cgraph.Node {
	id := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(id)
	node.SetLabel(e.Id)

	return node
}

func (l *SemControlListener) IdentifiantEnter(ctx parser.IIdentifiantContext) {
	// Nothing to do
}

func (l *SemControlListener) IdentifiantExit(ctx parser.IIdentifiantContext) {
	identifiant := &Identifiant{
		Id:  ctx.GetText(),
		Ctx: ctx,
	}

	t, err := l.Slt.GetSymbolType(ctx.GetText())

	if err != nil {
		// TODO: Handle error
	}

	identifiant.Type = t

	l.PushAst(identifiant)
}