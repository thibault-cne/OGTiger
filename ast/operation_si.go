package ast

import (
	"fmt"
	"ogtiger/logger"
	"ogtiger/parser"
	"ogtiger/slt"
	"ogtiger/ttype"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/goccy/go-graphviz/cgraph"
)

type OperationSi struct {
	Cond Ast
	Then Ast
	Else Ast
	Slt  *slt.SymbolTable
	Ctx  parser.IOperationSiContext
	Type *ttype.TigerType
}

func (e *OperationSi) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	condCtx := e.Cond.VisitSemControl(e.Slt, L)
	e.Then.VisitSemControl(e.Slt, L)

	if !e.Cond.ReturnType().Equals(ttype.NewTigerType(ttype.Int)) {
		L.NewSemanticError("Condition of the if is not an integer", condCtx)
	}
	
	if e.Else != nil {
		e.Else.VisitSemControl(e.Slt, L)

		if !e.Then.ReturnType().Equals(e.Else.ReturnType()) {
			L.NewSemanticError(logger.ErrorWrongTypeInIfElse, e.Ctx, e.Then.ReturnType(), e.Else.ReturnType())
		}
	}

	return e.Ctx
}

func (e *OperationSi) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *OperationSi) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("Si")

	cond := e.Cond.Draw(g)
	g.CreateEdge("Cond", node, cond)

	then := e.Then.Draw(g)
	g.CreateEdge("Then", node, then)

	if e.Else != nil {
		Else := e.Else.Draw(g)
		g.CreateEdge("Else", node, Else)
	}

	return node
}

func (l *AstCreatorListener) OperationSiEnter(ctx parser.IOperationSiContext) {
	// Creata a new region
	l.Slt = l.Slt.CreateRegion()
}

func (l *AstCreatorListener) OperationSiExit(ctx parser.IOperationSiContext) {
	OperationSi := &OperationSi{
		Ctx:  ctx,
		Type: ttype.NewTigerType(ttype.NoReturn),
	}

	if ctx.GetChildCount() == 6 {
		OperationSi.Else = l.PopAst()
	}

	OperationSi.Then = l.PopAst()
	OperationSi.Cond = l.PopAst()

	// Leave the region
	l.Slt = l.Slt.Parent
	OperationSi.Slt = l.Slt

	// Set the type of the operation
	OperationSi.Type = OperationSi.Then.ReturnType()

	// TODO: Check the type of the else

	// Push the new element on the stack
	l.PushAst(OperationSi)
}
