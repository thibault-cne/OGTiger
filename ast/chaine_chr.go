package ast

import (
	"fmt"
	"ogtiger/asm"
	"ogtiger/logger"
	"ogtiger/parser"
	"ogtiger/slt"
	"ogtiger/ttype"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/goccy/go-graphviz/cgraph"
)

type ChaineChr struct {
	Valeur string
	Ctx    parser.ChaineChrContext
	Type   *ttype.TigerType
}

func (e *ChaineChr) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	return &e.Ctx
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

func (l *AstCreatorListener) ChaineChrEnter(ctx parser.ChaineChrContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) ChaineChrExit(ctx parser.ChaineChrContext) {
	// Get back the last element of the stack
	chainChr := &ChaineChr{
		Ctx:  ctx,
		Type: ttype.NewTigerType(ttype.String),
	}

	chainChr.Valeur = ctx.GetChild(0).(*antlr.TerminalNodeImpl).GetText()

	l.PushAst(chainChr)
}

func (e *ChaineChr) EnterAsm(writer *asm.AssemblyWriter, slt *slt.SymbolTable) {
	defer e.ExitAsm(writer, slt)
}

func (e *ChaineChr) ExitAsm(writer *asm.AssemblyWriter, slt *slt.SymbolTable) {
	// Nothing to do
}
