package tokenizer

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	TK_ILLEGAL = "TK_ILLEGAL"
	EOF        = "EOF"

	// keywords

	TK_CREATE = "TK_CREATE"
	TK_FROM   = "TK_FROM"
	TK_INSERT = "TK_INSERT"
	TK_INTO   = "TK_INTO"
	TK_SELECT = "TK_SELECT"
	TK_TABLE  = "TK_TABLE"
	TK_WHERE  = "TK_WHERE"

	// identifier

	TK_ID = "TK_ID"

	// types

	TK_INTEGER = "TK_INTEGER"

	// operators
	TK_SEMI = "TK_SEMI"
	TK_LP   = "TK_LP"
	TK_RP   = "TK_RP"
	TK_STAR = "TK_STAR"
)

var keywords = map[string]TokenType{
	"CREATE":  TK_CREATE,
	"FROM":    TK_FROM,
	"INSERT":  TK_INSERT,
	"INTEGER": TK_INTEGER,
	"INTO":    TK_INTO,
	"SELECT":  TK_SELECT,
	"TABLE":   TK_TABLE,
	"WHERE":   TK_WHERE,
}

func LookupKeyword(identifier string) TokenType {
	if token, ok := keywords[identifier]; ok {
		return token
	}
	return TK_ID
}
