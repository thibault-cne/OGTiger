package ast

import (
	"fmt"
	"ogtiger/parser"
)

type Identifiant struct {
	Id  string
	Ctx parser.IIdentifiantContext
}

func (e *Identifiant) Display() string {
	return fmt.Sprintf(" identifiant %s", e.Id)
}

func (e *Identifiant) Draw(prefix string) {
	// TODO: Draw the AST
}

func (l *AstCreatorListener) IdentifiantEnter(ctx parser.IIdentifiantContext) {
	// Nothing to do
}

func (l *AstCreatorListener) IdentifiantExit(ctx parser.IIdentifiantContext) {
	identifiant := &Identifiant{
		Id:  ctx.GetText(),
		Ctx: ctx,
	}

	l.PushAst(identifiant)
}
