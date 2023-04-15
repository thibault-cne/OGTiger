package ast

import (
	"fmt"
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/goccy/go-graphviz/cgraph"
)

type ChaineChr struct {
	Valeur string
	Ctx    parser.ChaineChrContext
	Type   *ttype.TigerType
}

func (e *ChaineChr) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *ChaineChr) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel(e.Valeur)

	return node
}

func (l *SemControlListener) ChaineChrEnter(ctx parser.ChaineChrContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *SemControlListener) ChaineChrExit(ctx parser.ChaineChrContext) {
	// Get back the last element of the stack
	chainChr := &ChaineChr{
		Ctx:  ctx,
		Type: ttype.NewTigerType(ttype.String),
	}

	chainChr.Valeur = ctx.GetChild(0).(*antlr.TerminalNodeImpl).GetText()

	l.PushAst(chainChr)
}
