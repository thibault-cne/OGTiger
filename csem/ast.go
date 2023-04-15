package ast

import (
	"bytes"
	"fmt"
	"log"
	"ogtiger/logger"
	"ogtiger/parser"
	"ogtiger/slt"
	"ogtiger/ttype"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

type SemControlListener struct {
	*parser.BasetigerVisitor

	// The symbol table
	Slt *slt.SymbolTable

	// The logger
	Logger *logger.StepLogger
}

func NewSemControlListener(slt *slt.SymbolTable, L *logger.StepLogger) *SemControlListener {
	return &SemControlListener{
		Slt:    slt,
		Logger: L,
	}
}

type Ast interface {
	Draw(g *cgraph.Graph) *cgraph.Node
	ReturnType() *ttype.TigerType
}

// Example
func (ast *SemControlListener) VisitTerminal(node antlr.TerminalNode) {
	// fmt.Printf("%v\n", node.GetText())
}

func (ast *SemControlListener) VisitErrorNode(node antlr.ErrorNode) {
	// fmt.Printf("%v\n", node.GetText())
}

func (s *SemControlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	switch c := ctx.(type) {
	case parser.IProgramContext:
		s.ProgramEnter(c)
	case parser.IExpressionContext:
		s.ExprEnter(c)
	case parser.IOperationOuContext:
		s.OperationOuEnter(c)
	case parser.IOperationEtContext:
		s.OperationEtEnter(c)
	case parser.IOperationComparaisonContext:
		s.OperationComparaisonEnter(c)
	case parser.IOperationAdditionContext:
		s.OperationAdditionEnter(c)
	case parser.IOperationMultiplicationContext:
		s.OperationMultiplicationEnter(c)
	case parser.ISequenceInstructionContext:
		s.SequenceInstructionEnter(c)
	case parser.IOperationNegationContext:
		s.OperationNegationEnter(c)
	case parser.IExpressionValeurContext:
		switch c := c.(type) {
		case *parser.ExpressionIdentifiantContext:
			s.ExpressionIdentifiantEnter(*c)
		case *parser.AppelFonctionContext:
			s.AppelFonctionEnter(*c)
		case *parser.ListeAccesContext:
			s.ListAccesEnter(*c)
		case *parser.InstanciationTypeContext:
			s.InstanciationTypeEnter(*c)
		}
	case parser.IOperationSiContext:
		s.OperationSiEnter(c)
	case parser.IOperationTantqueContext:
		s.OperationTantQueEnter(c)
	case parser.IOperationBoucleContext:
		s.OperationBoucleEnter(c)
	case parser.IDefinitionContext:
		s.DefinitionEnter(c)
	case parser.IDeclarationTypeContext:
		switch c := c.(type) {
		case *parser.DeclarationTypeClassiqueContext:
			s.DeclarationTypeClassiqueEnter(*c)
		case *parser.DeclarationArrayTypeContext:
			s.DeclarationArrayTypeEnter(*c)
		case *parser.DeclarationRecordTypeContext:
			s.DeclarationRecordTypeEnter(*c)
		}
	case parser.IDeclarationChampContext:
		s.DeclarationChampEnter(c)
	case parser.IDeclarationFonctionContext:
		s.DeclarationFontionEnter(c)
	case parser.IDeclarationValeurContext:
		s.DeclarationValeurEnter(c)
	case parser.IConstantesContext:
		switch c := c.(type) {
		case *parser.ChaineChrContext:
			s.ChaineChrEnter(*c)
		case *parser.EntierContext:
			s.IntegerEnter(*c)
		case *parser.NilContext:
			s.NilEnter(*c)
		case *parser.BreakContext:
			s.BreakEnter(*c)
		}
	case parser.IIdentifiantContext:
		s.IdentifiantEnter(c)
	default:
		break
	}
}

func (s *SemControlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	switch c := ctx.(type) {
	case parser.IProgramContext:
		s.ProgramExit(c)
	case parser.IExpressionContext:
		s.ExprExit(c)
	case parser.IOperationOuContext:
		s.OperationOuExit(c)
	case parser.IOperationEtContext:
		s.OperationEtExit(c)
	case parser.IOperationComparaisonContext:
		s.OperationComparaisonExit(c)
	case parser.IOperationAdditionContext:
		s.OperationAdditionExit(c)
	case parser.IOperationMultiplicationContext:
		s.OperationMultiplicationExit(c)
	case parser.ISequenceInstructionContext:
		s.SequenceInstructionExit(c)
	case parser.IOperationNegationContext:
		s.OperationNegationExit(c)
	case parser.IExpressionValeurContext:
		switch c := c.(type) {
		case *parser.ExpressionIdentifiantContext:
			s.ExpressionIdentifiantExit(*c)
		case *parser.AppelFonctionContext:
			s.AppelFonctionExit(*c)
		case *parser.ListeAccesContext:
			s.ListAccesExit(*c)
		case *parser.InstanciationTypeContext:
			s.InstanciationTypeExit(*c)
		}
	case parser.IOperationSiContext:
		s.OperationSiExit(c)
	case parser.IOperationTantqueContext:
		s.OperationTantQueExit(c)
	case parser.IOperationBoucleContext:
		s.OperationBoucleExit(c)
	case parser.IDefinitionContext:
		s.DefinitionExit(c)
	case parser.IDeclarationTypeContext:
		switch c := c.(type) {
		case *parser.DeclarationTypeClassiqueContext:
			s.DeclarationTypeClassiqueExit(*c)
		case *parser.DeclarationArrayTypeContext:
			s.DeclarationArrayTypeExit(*c)
		case *parser.DeclarationRecordTypeContext:
			s.DeclarationRecordTypeExit(*c)
		}
	case parser.IDeclarationChampContext:
		s.DeclarationChampExit(c)
	case parser.IDeclarationFonctionContext:
		s.DeclarationFontionExit(c)
	case parser.IDeclarationValeurContext:
		s.DeclarationValeurExit(c)
	case parser.IConstantesContext:
		switch c := c.(type) {
		case *parser.ChaineChrContext:
			s.ChaineChrExit(*c)
		case *parser.EntierContext:
			s.IntegerExit(*c)
		case *parser.NilContext:
			s.NilExit(*c)
		case *parser.BreakContext:
			s.BreakExit(*c)
		}
	case parser.IIdentifiantContext:
		s.IdentifiantExit(c)
	default:
		break
	}
}

func (s *SemControlListener) DisplayAST(filename string) {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		log.Fatal(err)
	}

	s.AstStack[0].Draw(graph)

	var buf bytes.Buffer
	if err := g.Render(graph, "dot", &buf); err != nil {
		log.Fatal(err)
	}

	path := fmt.Sprintf("output/%s.png", filename)
	// Write to file
	if err := g.RenderFilename(graph, graphviz.PNG, path); err != nil {
		log.Fatal(err)
	}
}
