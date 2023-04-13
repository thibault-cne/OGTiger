package ast

import "ogtiger/parser"

type OperationEt struct {
	Left  Ast
	Right []Ast
}

func (e OperationEt) Display() string {
	return " et"
}

func (l *AstCreatorListener) OperationEtEnter(ctx parser.IOperationEtContext) {
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func (l *AstCreatorListener) OperationEtExit(ctx parser.IOperationEtContext) {
	operationEt := &OperationEt{}

	if ctx.GetChildCount() == 1 {
		return
	}

	operationEt.Left = l.PopAst()

	// Get the other exprEt
	for i := 0; i < (ctx.GetChildCount()-1)/2; i++ {
		operationEt.Right = append(operationEt.Right, l.PopAst())
	}

	l.PushAst(operationEt)
}
