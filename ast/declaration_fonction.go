package ast

import (
	"fmt"
	"ogtiger/logger"
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type DeclarationFontion struct {
	Id    Ast
	Args  []Ast
	FType Ast
	Expr  Ast
	Ctx   parser.IDeclarationFonctionContext
	Type  *ttype.TigerType
}

func (e *DeclarationFontion) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *DeclarationFontion) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("DeclarationFontion")

	id := e.Id.Draw(g)
	g.CreateEdge("Id", node, id)

	for _, arg := range e.Args {
		argNode := arg.Draw(g)
		g.CreateEdge("Arg", node, argNode)
	}

	if e.FType != nil {
		typeNode := e.FType.Draw(g)
		g.CreateEdge("Type", node, typeNode)
	}

	exprNode := e.Expr.Draw(g)
	g.CreateEdge("Expr", node, exprNode)

	return node
}

func (l *AstCreatorListener) DeclarationFontionEnter(ctx parser.IDeclarationFonctionContext) {
	// Nothing to do
}

func (l *AstCreatorListener) DeclarationFontionExit(ctx parser.IDeclarationFonctionContext) {
	declarationFontion := &DeclarationFontion{Ctx: ctx}

	declarationFontion.Expr = l.PopAst()

	// Pop the type if it exists
	if len(ctx.AllIdentifiant()) > 1 {
		declarationFontion.FType = l.PopAst()
	}

	// Pop all args
	args := []*ttype.FunctionParameter{}
	for i := 0; i < len(ctx.AllDeclarationChamp()); i++ {
		// Prepend the arg
		declarationFontion.Args = append([]Ast{l.PopAst()}, declarationFontion.Args...)
		args = append(args, &ttype.FunctionParameter{
			Name: declarationFontion.Args[i].(*DeclarationChamp).Left.(*Identifiant).Id,
			Type: declarationFontion.Args[i].(*DeclarationChamp).Right.(*Identifiant).ReturnType(),
		})
	}

	declarationFontion.Id = l.PopAst()

	// Verify that the Id is not already defined
	if _, err := l.Slt.GetSymbol(declarationFontion.Id.(*Identifiant).Id); err == nil {
		l.Logger.NewSemanticError(logger.ErrorIdIsAlreadyDefinedInScope, ctx, declarationFontion.Id.(*Identifiant).Id)
	}

	l.PushAst(declarationFontion)
}
