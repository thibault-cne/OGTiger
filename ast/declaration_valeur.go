package ast

import "ogtiger/parser"

type DeclarationValeur struct {
	Id Ast
	Type Ast
	Expr Ast
	Ctx  parser.IDeclarationValeurContext
}

func (e DeclarationValeur) Display() string {
	return " declarationValeur"
}

func (l *AstCreatorListener) DeclarationValeurEnter(ctx parser.IDeclarationValeurContext) {
	// Nothing to do
}

func (l *AstCreatorListener) DeclarationValeurExit(ctx parser.IDeclarationValeurContext) {
	declarationValeur := &DeclarationValeur{
		Ctx: ctx,
	}

	declarationValeur.Id = l.PopAst()

	if len(ctx.AllIdentifiant()) > 1 {
		declarationValeur.Type = l.PopAst()
	}

	declarationValeur.Expr = l.PopAst()

	l.PushAst(declarationValeur)
}