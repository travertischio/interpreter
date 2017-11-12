package parser

import (
	"github.com/travertischio/interpreter/ast"
	"github.com/travertischio/interpreter/lexer"
	"github.com/travertischio/interpreter/token"
)

type Parser struct {
	l *lexer.lexer

	curToken  token.Token
	peekToken token.Token
}

func new(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens for curToken and peekToken
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
