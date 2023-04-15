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

type OperationAddition struct {
	Left  Ast
	Right Ast
	Ctx   parser.IOperationAdditionContext
	Type  *ttype.TigerType
}

func (e *OperationAddition) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	// TODO: Fill this
	return e.Ctx
}

type OperationSoustraction struct {
	Left  Ast
	Right Ast
	Ctx   parser.IOperationAdditionContext
	Type  *ttype.TigerType
}

func (e *OperationSoustraction) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	// TODO: Fill this
	return e.Ctx
}

func (e *OperationAddition) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *OperationSoustraction) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *OperationAddition) Draw(g *cgraph.Graph) *cgraph.Node {
	id := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(id)
	node.SetLabel("+")

	left := e.Left.Draw(g)
	g.CreateEdge("Left", node, left)
	right := e.Right.Draw(g)
	g.CreateEdge("Right", node, right)

	return node
}

func (e *OperationSoustraction) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("-")

	left := e.Left.Draw(g)
	g.CreateEdge("Left", node, left)
	right := e.Right.Draw(g)
	g.CreateEdge("Right", node, right)

	return node
}

func (l *AstCreatorListener) OperationAdditionEnter(ctx parser.IOperationAdditionContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) OperationAdditionExit(ctx parser.IOperationAdditionContext) {
	if ctx.GetChildCount() == 1 {
		return
	}

	// Get back the elements needed from the stack
	elements := make([]Ast, 0)

	for i := 0; i < (ctx.GetChildCount()+1)/2; i++ {
		elements = append(elements, l.PopAst())
	}

	node := elements[len(elements)-1]
	elements = elements[:len(elements)-1]

	// TODO: Check if the type is correct

	// Get minus and plus and term number
	for i := 0; 2*i < (ctx.GetChildCount() - 1); i++ {
		switch ctx.GetChild(2*i + 1).(*antlr.TerminalNodeImpl).GetText() {
		case "+":
			temp := &OperationAddition{
				Ctx:   ctx,
				Left:  node,
				Right: elements[len(elements)-1],
				Type:  ttype.NewTigerType(ttype.Int),
			}
			node = temp
		case "-":
			temp := &OperationSoustraction{
				Ctx:   ctx,
				Left:  node,
				Right: elements[len(elements)-1],
				Type:  ttype.NewTigerType(ttype.Int),
			}
			node = temp
		}

		elements = elements[:len(elements)-1]
	}

	l.PushAst(node)
}
