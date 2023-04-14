package ast

import (
	"ogtiger/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/goccy/go-graphviz/cgraph"
)

type OperationAddition struct {
	Left  Ast
	Right []*OperationAdditionFD
	Ctx   parser.IOperationAdditionContext
}

type OperationAdditionFD struct {
	Op    string
	Right Ast
}

func (e *OperationAddition) Display() string {
	return " addition"
}

func (e *OperationAddition) Draw(prefix string, g *cgraph.Graph) {
	// TODO: Draw the AST
}

func (l *AstCreatorListener) OperationAdditionEnter(ctx parser.IOperationAdditionContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) OperationAdditionExit(ctx parser.IOperationAdditionContext) {
	// Get back the last element of the stack
	opAddition := &OperationAddition{
		Ctx: ctx,
	}

	if ctx.GetChildCount() == 1 {
		return
	}

	opAddition.Left = l.PopAst()

	// Get minus and plus and term number
	for i := 0; i < (ctx.GetChildCount()-1)/2; i++ {
		right := &OperationAdditionFD{}

		right.Op = ctx.GetChild(2*i + 1).(*antlr.TerminalNodeImpl).GetText()
		right.Right = l.PopAst()

		opAddition.Right = append(opAddition.Right, right)
	}

	l.PushAst(opAddition)
}
