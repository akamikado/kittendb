package parser

import "kittendb/tokenizer"

type Parser struct {
	tokenizer *tokenizer.Tokenizer

	pos   tokenizer.Pos
	token tokenizer.Token
	full  bool
}

func New(t *tokenizer.Tokenizer) *Parser {
	p := &Parser{tokenizer: t}
	return p
}

func (p *Parser) Peek() tokenizer.Token {
	return p.tokenizer.GetToken()
}

func (p *Parser) EvalSQL(input string) *Query {
	stmts := &Query{}
	stmts.Statements = []Statement{}
	for !p.full {
		stmt := p.ParseStatement()
		stmts.Statements = append(stmts.Statements, stmt)
	}

	return stmts
}

func (p *Parser) ParseStatement() Statement {}

func (p *Parser) ParseSelectStatement() SelectStatement {}
