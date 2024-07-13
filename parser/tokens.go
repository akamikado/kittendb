package parser

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	TK_ILLEGAL = "TK_ILLEGAL"
	EOF        = "EOF"

	// keywords

	TK_CREATE   = "TK_CREATE"
	TK_DISTINCT = "TK_DISTINCT"
	TK_FROM     = "TK_FROM"
	TK_INSERT   = "TK_INSERT"
	TK_INTO     = "TK_INTO"
	TK_NULL     = "TK_NULL"
	TK_SELECT   = "TK_SELECT"
	TK_TABLE    = "TK_TABLE"
	TK_VALUES   = "TK_VALUES"
	TK_WHERE    = "TK_WHERE"

	// identifier

	TK_ID = "TK_ID"

	// types

	TK_INTEGER = "TK_INTEGER"

	// operators

	TK_SEMI  = "TK_SEMI"
	TK_LP    = "TK_LP"
	TK_RP    = "TK_RP"
	TK_STAR  = "TK_STAR"
	TK_COMMA = "TK_COMMA"
	TK_PLUS  = "TK_PLUS"
	TK_MINUS = "TK_MINUS"
	TK_SLASH = "TK_SLASH"
	TK_GT    = "TK_GT"
	TK_LE    = "TK_LE"
	TK_LT    = "TK_LT"
	TK_GE    = "TK_GE"
	TK_EQ    = "TK_EQ"
	TK_NE    = "TK_NE"

	TK_OR  = "TK_OR"
	TK_AND = "TK_AND"
	TK_NOT = "TK_NOT"
)

var keywords = map[string]TokenType{
	"CREATE":   TK_CREATE,
	"DISTINCT": TK_DISTINCT,
	"FROM":     TK_FROM,
	"INSERT":   TK_INSERT,
	"INTEGER":  TK_INTEGER,
	"INTO":     TK_INTO,
	"NULL":     TK_NULL,
	"SELECT":   TK_SELECT,
	"TABLE":    TK_TABLE,
	"WHERE":    TK_WHERE,
}

func LookupKeyword(identifier string) TokenType {
	if token, ok := keywords[identifier]; ok {
		return token
	}
	return TK_ID
}

const (
	LowestPrec = 0
)

func (op TokenType) Precedence() int {
	switch op {
	case TK_OR:
		return 1
	case TK_AND:
		return 2
	case TK_NOT:
		return 3
	case TK_NE, TK_EQ:
		return 4
	case TK_GT, TK_LE, TK_LT, TK_GE:
		return 5
	case TK_PLUS, TK_MINUS:
		return 8
	case TK_STAR, TK_SLASH:
		return 9
	}
	return LowestPrec
}
