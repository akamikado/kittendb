package parser

import "kittendb/tokenizer"

type Parser struct {
	tokenizer *tokenizer.Tokenizer

	pos   tokenizer.Pos
	token tokenizer.Token
	full  bool
}

func New(input string) *Parser {
	t := tokenizer.New(input)
	p := &Parser{tokenizer: t}
	return p
}

func (p *Parser) Peek() tokenizer.Token {
	token := p.tokenizer.PeekToken()
	return token
}

func (p *Parser) NextToken() {
	p.token = p.tokenizer.GetToken()
	if p.token.Type == tokenizer.EOF {
		p.full = true
	}
	p.pos = p.tokenizer.GetPos()
}

func (p *Parser) EvalSQL(input string) (*Query, error) {
	stmts := &Query{}
	stmts.Statements = []Statement{}
	for !p.full {
		stmt, err := p.ParseStatement()
		if err != nil {
			return nil, err
		}
		stmts.Statements = append(stmts.Statements, stmt)
	}

	return stmts, nil
}

func (p *Parser) ParseStatement() (Statement, error) {
	switch p.Peek().Type {
	case tokenizer.TK_SELECT:
		return p.ParseSelectStatement()
	default:
		return nil, nil
	}
}

func (p *Parser) ParseSelectStatement() (*SelectStatement, error) {
	stmt := &SelectStatement{Select: p.pos}

	return stmt, nil
}
