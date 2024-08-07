package parser

import (
	"github.com/akamikado/kittendb/tokenizer"
)

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
	p.token, p.pos = p.tokenizer.GetToken()
	if p.token.Type == tokenizer.EOF {
		p.full = true
	}
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
	case tokenizer.TK_INSERT:
		return p.ParseInsertStatement()
	default:
		// TODO: errors
		return nil, nil
	}
}

func (p *Parser) ParseSelectStatement() (*SelectStatement, error) {
	p.NextToken()
	stmt := &SelectStatement{Select: p.pos}

	p.NextToken()
	if p.token.Type == tokenizer.TK_DISTINCT {
		stmt.Distinct = p.pos
		p.NextToken()
	}

	if p.token.Type == tokenizer.TK_FROM {
		// TODO: errors
		return nil, nil
	}

	for p.token.Type != tokenizer.TK_FROM {
		if p.token.Type == tokenizer.TK_COMMA {
			p.NextToken()
			continue
		}

		rcs, err := p.ParseResultColumn()
		if err != nil {
			return nil, err
		}
		stmt.Columns = append(stmt.Columns, rcs)
		p.NextToken()
	}

	if p.token.Type == tokenizer.TK_FROM {
		stmt.From = p.pos
		p.NextToken()
	} else {
		// TODO: errors
		return nil, nil
	}

	source, err := p.ParseSource()
	if err != nil {
		return nil, err
	}
	stmt.Source = source

	return stmt, nil
}

func (p *Parser) ParseResultColumn() (*ResultColumn, error) {
	switch p.token.Type {
	case tokenizer.TK_STAR:
		star := &ResultColumn{Star: p.pos}
		return star, nil

	case tokenizer.TK_ID:
		word, err := p.ParseIdentifier()
		if err != nil {
			return nil, err
		}
		ident := &ResultColumn{Expr: word}
		return ident, nil

	default:
		// TODO: errors
		return nil, nil
	}
}

func (p *Parser) ParseIdentifier() (*Identifier, error) {
	ident := &Identifier{NamePos: p.pos, Name: p.token.Literal}
	return ident, nil
}

func (p *Parser) ParseSource() (Source, error) {
	src, err := p.ParseUnarySource()
	if err != nil {
		return nil, err
	}
	return src, nil
}

func (p *Parser) ParseUnarySource() (Source, error) {
	switch p.token.Type {
	case tokenizer.TK_ID:
		return p.ParseTableName()
	default:
		// TODO: errors
		return nil, nil
	}
}

func (p *Parser) ParseTableName() (Source, error) {
	table := &Identifier{NamePos: p.pos, Name: p.token.Literal}
	src := &TableName{Name: table}
	return src, nil
}

func (p *Parser) ParseInsertStatement() (*InsertStatement, error) {
	var err error

	p.NextToken()
	stmt := &InsertStatement{Insert: p.pos}

	p.NextToken()
	if p.token.Type == tokenizer.TK_INTO {
		stmt.Into = p.pos
	}

	p.NextToken()
	if stmt.Table, err = p.ParseIdentifier(); err != nil {
		return nil, err
	}

	p.NextToken()
	if p.token.Type == tokenizer.TK_LP {
		stmt.ColumnsLparen = p.pos
		p.NextToken()
		for {
			if p.token.Type == tokenizer.TK_RP {
				break
			}
			column, err := p.ParseIdentifier()
			if err != nil {
				return nil, err
			}
			stmt.Columns = append(stmt.Columns, column)

			p.NextToken()
			if p.token.Type == tokenizer.TK_RP {
				break
			} else if p.token.Type != tokenizer.TK_COMMA {
				// TODO: errors
				return nil, nil
			}
			p.NextToken()
		}
		stmt.ColumnsRparen = p.pos

		p.NextToken()
		switch p.token.Type {
		case tokenizer.TK_VALUES:
			stmt.Values = p.pos
			for {
				var list ExprList
				p.NextToken()
				if p.token.Type != tokenizer.TK_LP {
					// TODO: errors
					return nil, nil
				}
				p.NextToken()
				for {
					expr, err := p.ParseExpression()
					if err != nil {
						return nil, err
					}
					list.Exprs = append(list.Exprs, expr)
					p.NextToken()
					if p.token.Type == tokenizer.TK_RP {
						break
					} else if p.token.Type != tokenizer.TK_COMMA {
						// TODO: errors
						return nil, nil
					}
					p.NextToken()
				}
				stmt.ValuesList = append(stmt.ValuesList, &list)

				if p.Peek().Type == tokenizer.TK_COMMA {
					break
				}
				p.NextToken()
			}

		case tokenizer.TK_SELECT:
			// TODO: implement

		default:
			// TODO: errors
			return nil, nil
		}
	}

	return stmt, nil
}

func (p *Parser) ParseExpression() (Expression, error) {
	return p.ParseBinaryExpression(tokenizer.LowestPrec + 1)
}

func (p *Parser) ParseBinaryExpression(prec int) (Expression, error) {
	x, err := p.ParseOperand()
	if err != nil {
		return nil, err
	}
	return x, nil
}

func (p *Parser) ParseOperand() (Expression, error) {
	switch p.token.Type {
	case tokenizer.TK_INTEGER:
		return &NumberLit{ValuePos: p.pos, Value: p.token.Literal}, nil
	case tokenizer.TK_NULL:
		return &NullLit{Pos: p.pos}, nil

	default:
		// TODO: errors
		return nil, nil
	}
}
