package parser

import (
	"bytes"
	"fmt"

	"github.com/akamikado/kittendb/tokenizer"
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

func (*NullLit) node() {}

func (*NumberLit) node() {}

func (*ResultColumn) node() {}

func (*SelectStatement) node() {}

func (*TableName) node() {}

func (*Type) node() {}

type Expression interface {
	Node
	expr()
}

func (*ExprList) expr() {}

func (*Identifier) expr() {}

func (*NullLit) expr() {}

func (*NumberLit) expr() {}

type Statement interface {
	Node
	stmt()
}

func (*InsertStatement) stmt() {}

func (*SelectStatement) stmt() {}

type Query struct {
	Statements []Statement
}

type ExprList struct {
	Lparen tokenizer.Pos
	Exprs  []Expression
	Rparen tokenizer.Pos
}

func (el *ExprList) String() string {
	var buf bytes.Buffer
	buf.WriteString("(")
	for i, expr := range el.Exprs {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(expr.String())
	}
	buf.WriteString(")")

	return buf.String()
}

type Identifier struct {
	NamePos tokenizer.Pos
	Name    string
}

func (i *Identifier) String() string {
	return i.Name
}

type NullLit struct {
	Pos tokenizer.Pos
}

func (nl *NullLit) String() string {
	return "NULL"
}

type NumberLit struct {
	ValuePos tokenizer.Pos
	Value    string
}

func (nl *NumberLit) String() string {
	return nl.Value
}

type Type struct {
	Name *Identifier
}

func (t *Type) String() string {
	return t.Name.Name
}

type ColumnDefinition struct {
	Name *Identifier
	Type *Type
}

func (cd *ColumnDefinition) String() string {
	var buf bytes.Buffer
	buf.WriteString(cd.Name.String())
	buf.WriteString(" ")
	buf.WriteString(cd.Type.String())
	return buf.String()
}

type TableName struct {
	Name *Identifier
}

func (tn *TableName) String() string {
	var buf bytes.Buffer
	buf.WriteString(tn.Name.String())
	return buf.String()
}

type ResultColumn struct {
	Star tokenizer.Pos
	Expr Expression
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

func (*TableName) source() {}

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

// CREATE TABLE Statement
type CreateTableStatement struct {
	Create tokenizer.Pos
	Table  tokenizer.Pos
	Name   *Identifier

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

// INSERT Statement
type InsertStatement struct {
	Insert tokenizer.Pos
	Into   tokenizer.Pos

	Table *Identifier

	ColumnsLparen tokenizer.Pos
	Columns       []*Identifier
	ColumnsRparen tokenizer.Pos

	Values     tokenizer.Pos
	ValuesList []*ExprList
}

func (is *InsertStatement) String() string {
	var buf bytes.Buffer
	buf.WriteString("INSERT INTO ")
	buf.WriteString(is.Table.String())
	buf.WriteString(" (")
	for _, col := range is.Columns {
		if col != is.Columns[0] {
			buf.WriteString(", ")
		}
		buf.WriteString(col.String())
	}
	buf.WriteString(") VALUES ")
	for _, expr := range is.ValuesList {
		if expr != is.ValuesList[0] {
			buf.WriteString(", ")
		}
		buf.WriteString(expr.String())
	}

	return buf.String()
}
