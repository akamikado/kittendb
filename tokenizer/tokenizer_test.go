package tokenizer

import (
	"reflect"
	"testing"
)

func TestTokenizer(t *testing.T) {
	tests := []struct {
		input               string
		expectedIdentifiers []Token
	}{
		{"123", []Token{{Type: TK_INTEGER, Literal: "123"}, {Type: EOF, Literal: ""}}},
		{"abc", []Token{{Type: TK_ID, Literal: "abc"}, {Type: EOF, Literal: ""}}},
		{"abc 123", []Token{{Type: TK_ID, Literal: "abc"}, {Type: TK_INTEGER, Literal: "123"}, {Type: EOF, Literal: ""}}},
		{"(abc 123)", []Token{{Type: TK_LP, Literal: "("}, {Type: TK_ID, Literal: "abc"}, {Type: TK_INTEGER, Literal: "123"}, {Type: TK_RP, Literal: ")"}, {Type: EOF, Literal: ""}}},
		{"SELECT * FROM abc;", []Token{{Type: TK_SELECT, Literal: "SELECT"}, {Type: TK_STAR, Literal: "*"}, {Type: TK_FROM, Literal: "FROM"}, {Type: TK_ID, Literal: "abc"}, {Type: TK_SEMI, Literal: ";"}, {Type: EOF, Literal: ""}}},
		{"CREATE TABLE abc (id INTEGER);", []Token{{Type: TK_CREATE, Literal: "CREATE"}, {Type: TK_TABLE, Literal: "TABLE"}, {Type: TK_ID, Literal: "abc"}, {Type: TK_LP, Literal: "("}, {Type: TK_ID, Literal: "id"}, {Type: TK_INTEGER, Literal: "INTEGER"}, {Type: TK_RP, Literal: ")"}, {Type: TK_SEMI, Literal: ";"}, {Type: EOF, Literal: ""}}},
	}

	for _, test := range tests {
		tokenizer := New([]byte(test.input))

		var tokens []Token

		tokens = append(tokens, tokenizer.GetToken())
		for tokens[len(tokens)-1].Type != EOF {
			tokens = append(tokens, tokenizer.GetToken())
		}

		if !reflect.DeepEqual(tokens, test.expectedIdentifiers) {
			t.Errorf("Expected %v, got %v", test.expectedIdentifiers, tokens)
		}

	}
}
