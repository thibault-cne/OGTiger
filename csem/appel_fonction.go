package ast

import (
	"fmt"
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type AppelFonction struct {
	Identifiant Ast
	Args        []Ast
	Ctx         parser.AppelFonctionContext
	Type        *ttype.TigerType
}

func (e *AppelFonction) ReturnType() *ttype.TigerType {
	return e.Type
}

func (a *AppelFonction) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", a)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("AppelFonction")

	id := a.Identifiant.Draw(g)
	g.CreateEdge("Id", node, id)

	nodeAId := fmt.Sprintf("NA%p", a.Args)
	nodeA, _ := g.CreateNode(nodeAId)
	nodeA.SetLabel("ArgFonction")

	g.CreateEdge("Args", node, nodeA)

	for _, arg := range a.Args {
		argNode := arg.Draw(g)
		g.CreateEdge("", nodeA, argNode)
	}

	return node
}

func (l *SemControlListener) AppelFonctionEnter(ctx parser.AppelFonctionContext) {
	// Nothing to do
}

func (l *SemControlListener) AppelFonctionExit(ctx parser.AppelFonctionContext) {
	appelFonction := &AppelFonction{
		Ctx: ctx,
	}

	// Get the args
	for i := 0; i < len(ctx.AllExpression()); i++ {
		// Pop the last element of the stack
		// Prepend it to the args
		appelFonction.Args = append([]Ast{l.PopAst()}, appelFonction.Args...)
	}

	// Get the identifiant
	appelFonction.Identifiant = l.PopAst()

	// Set the type
	appelFonction.Type = appelFonction.Identifiant.ReturnType()

	// Push the new element on the stack
	l.PushAst(appelFonction)
}