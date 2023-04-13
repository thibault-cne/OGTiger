package ast

import "ogtiger/parser"

type ExpressionIdentifiant struct{}

func (e ExpressionIdentifiant) Display() string {
	return " expressionValeur"
}

func (l *AstCreatorListener) ExpressionIdentifiantEnter(ctx parser.ExpressionIdentifiantContext) {
	// Nothing to do
}

func (l *AstCreatorListener) ExpressionIdentifiantExit(ctx parser.ExpressionIdentifiantContext) {
	// Nothing to do
}
