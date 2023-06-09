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

type DeclarationFontion struct {
	Id    Ast
	Args  []Ast
	FType Ast
	Expr  Ast
	Ctx   parser.IDeclarationFonctionContext
	Slt   *slt.SymbolTable
	Type  *ttype.TigerType
}

func (e *DeclarationFontion) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	e.Id.VisitSemControl(e.Slt, L)
	for _, arg := range e.Args {
		arg.VisitSemControl(e.Slt, L)
	}
	if e.FType != nil {
		e.FType.VisitSemControl(e.Slt, L)

		// Verify that the type is defined
		if _, err := slt.GetSymbol(e.FType.(*Identifiant).Id); err != nil {
			L.NewSemanticError(logger.ErrorTypeIsNotDefined, e.FType.(*Identifiant).Ctx, e.FType.(*Identifiant).Id)
		}

		// Verify that the type is the same as the return type of the function
		if !e.FType.(*Identifiant).ReturnType().Equals(e.Expr.ReturnType()) {
			L.NewSemanticError(logger.ErrorFunctionMismatchedTypes, e.FType.(*Identifiant).Ctx, e.Id.(*Identifiant).Id, e.FType.ReturnType(), e.Expr.ReturnType())
		}
	}
	e.Expr.VisitSemControl(e.Slt, L)
	return e.Ctx
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
	// Create a TDS for the function
	l.Slt = l.Slt.CreateRegion()
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
		a := l.PopAst()
		declarationFontion.Args = append([]Ast{a}, declarationFontion.Args...)
		args = append(args, &ttype.FunctionParameter{
			Name: a.(*DeclarationChamp).Left.(*Identifiant).Id,
			Type: a.(*DeclarationChamp).Right.(*Identifiant).ReturnType(),
		})
	}

	// Set the type of the function
	if declarationFontion.FType == nil {
		declarationFontion.Type = ttype.NewFunctionType(
			declarationFontion.Expr.ReturnType(),
			args,
		)
	} else {
		declarationFontion.Type = ttype.NewFunctionType(
			declarationFontion.FType.ReturnType(),
			args,
		)
	}

	declarationFontion.Id = l.PopAst()

	// Verify that the Id is not already defined
	if _, err := l.Slt.GetSymbol(declarationFontion.Id.(*Identifiant).Id); err == nil {
		l.Logger.NewSemanticError(logger.ErrorIdIsAlreadyDefinedInScope, declarationFontion.Id.(*Identifiant).Ctx, declarationFontion.Id.(*Identifiant).Id)
	}

	// Add the parameters to the TDS
	for i, arg := range args {
		a := &slt.Symbol{
			Name:   arg.Name,
			Type:   arg.Type,
			Offset: (len(args)-i-1)*4 + 12,
		}

		l.Slt.AddSymbol(arg.Name, a)
	}

	f := &slt.Symbol{
		Name: declarationFontion.Id.(*Identifiant).Id,
		Type: &ttype.TigerType{
			ID:         ttype.Function,
			Parameters: args,
			ReturnType: declarationFontion.ReturnType(),
		},
		Offset:      0,
		SymbolTable: l.Slt,
	}

	l.Slt = l.Slt.Parent
	declarationFontion.Slt = l.Slt

	// Add the function to the TDS
	l.Slt.AddSymbol(declarationFontion.Id.(*Identifiant).Id, f)

	l.PushAst(declarationFontion)
}

func (e *DeclarationFontion) EnterAsm(writer *asm.AssemblyWriter, slt *slt.SymbolTable) {
	defer e.ExitAsm(writer, slt)
}

func (e *DeclarationFontion) ExitAsm(writer *asm.AssemblyWriter, slt *slt.SymbolTable) {
	// Nothing to do
}
