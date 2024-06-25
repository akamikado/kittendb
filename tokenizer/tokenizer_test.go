package tokenizer

import (
	"reflect"
	"strings"
	"testing"
)

func TestTokenizer(t *testing.T) {
	tests := []struct {
		input               []byte
		expectedIdentifiers []string
	}{
		{[]byte("123"), []string{TK_INTEGER, EOF}},
		{[]byte("abc"), []string{TK_ID, EOF}},
		{[]byte("SELECT * FROM abc"), []string{TK_SELECT, TK_STAR, TK_FROM, TK_ID, EOF}},
		{[]byte("CREATE TABLE abc (id INTEGER);"), []string{TK_CREATE, TK_TABLE, TK_ID, TK_LP, TK_ID, TK_INTEGER, TK_RP, TK_SEMI, EOF}},
	}

	for _, test := range tests {
		tokenizer := New([]byte(test.input))

		var tokens []Token

		tokens = append(tokens, tokenizer.GetToken())
		for tokens[len(tokens)-1].Type != EOF {
			tokens = append(tokens, tokenizer.GetToken())
		}

		var tokenTypes []string
		for _, token := range tokens {
			tokenTypes = append(tokenTypes, string(token.Type))
		}

		if !reflect.DeepEqual(tokenTypes, test.expectedIdentifiers) {
			t.Fatalf("expected (%s) got (%s)", strings.Join(test.expectedIdentifiers, " "), strings.Join(tokenTypes, " "))
		}
	}
}
