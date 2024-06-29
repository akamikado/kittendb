package parser

import (
	"bytes"
	"fmt"

	"kittendb/tokenizer"
)

type Node interface {
	node()
	fmt.Stringer
}

func (*ColumnDefinition) node() {}

func (*CreateTableStatement) node() {}

func (*ExprList) node() {}

func (*Identifier) node() {}

func (*InsertStatement) node() {}

func (*ResultColumn) node() {}

func (*SelectStatement) node() {}

func (*Type) node() {}

type Expression interface {
	Node
	expr()
}

func (*Identifier) expr() {}

func (*ExprList) expr() {}

type Statement interface {
	Node
	stmt()
}

func (*SelectStatement) stmt() {}

type ExprList struct {
	Lparen tokenizer.Pos
	Exprs  []Expression
	Rparen tokenizer.Pos
}

type Identifier struct {
	NamePos tokenizer.Pos //identifier position
	Name    string        // Identifier name
}

func (i *Identifier) String() string {
	return i.Name
}

type Type struct {
	Name *Identifier
}

type ColumnDefinition struct {
	Name *Identifier
	Type *Type
}

func (t *Type) String() string {
	return t.Name.Name
}

func (cd *ColumnDefinition) String() string {
	var buf bytes.Buffer
	buf.WriteString(cd.Name.String())
	buf.WriteString(" ")
	buf.WriteString(cd.Type.String())
	return buf.String()
}

type ResultColumn struct {
	Star tokenizer.Pos // position of *
	Expr Expression    // column expression
}

func (rc *ResultColumn) String() string {
	if rc.Star.IsValid() {
		return "*"
	}
	return rc.Expr.String()
}

// Source is table
type Source interface {
	Node
	source()
}

// SELECT Statement
type SelectStatement struct {
	Select   tokenizer.Pos
	Distinct tokenizer.Pos
	Columns  []*ResultColumn

	From   tokenizer.Pos
	Source Source

	Where     tokenizer.Pos
	WhereExpr Expression
}

func (s *SelectStatement) String() string {
	var buf bytes.Buffer

	buf.WriteString("SELECT ")
	if s.Distinct.IsValid() {
		buf.WriteString("DISTINCT ")
	}

	for i, col := range s.Columns {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(col.String())
	}

	if s.Source != nil {
		fmt.Fprintf(&buf, " FROM %s", s.Source.String())
	}

	if s.WhereExpr != nil {
		fmt.Fprintf(&buf, " WHERE %s", s.WhereExpr.String())
	}

	return buf.String()
}

type CreateTableStatement struct {
	Create tokenizer.Pos // position of CREATE keyword
	Table  tokenizer.Pos // position of TABLE keyword
	Name   *Identifier   // toble name

	Lparen  tokenizer.Pos
	Columns []*ColumnDefinition
	Rparen  tokenizer.Pos
}

func (cts *CreateTableStatement) String() string {
	var buf bytes.Buffer
	buf.WriteString("CREATE TABLE ")
	buf.WriteString(cts.Name.String())

	buf.WriteString(" (")

	for i := range cts.Columns {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(cts.Columns[i].String())
	}

	buf.WriteString(" )")

	return cts.String()
}

type InsertStatement struct {
	Insert tokenizer.Pos // position of INSERT keyword
	Table  *Identifier   // table name

	ColumnsLparen tokenizer.Pos // position of column list left paren
	Columns       []*Identifier // columns list
	ColumnsRparen tokenizer.Pos // position of column list right paren

	Values     tokenizer.Pos // position of VALUES keyword
	ValuesList []*ExprList   // list of list of values
}
