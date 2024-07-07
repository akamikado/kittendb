package tokenizer

import (
	"reflect"
	"testing"
)

func TestTokenizer(t *testing.T) {
	tests := []struct {
		input               string
		expectedIdentifiers []Token
		positions           []Pos
	}{
		{"123", []Token{{Type: TK_INTEGER, Literal: "123"}, {Type: EOF, Literal: ""}}, []Pos{{offset: 0}, {offset: 3}}},
		{"abc", []Token{{Type: TK_ID, Literal: "abc"}, {Type: EOF, Literal: ""}}, []Pos{{offset: 0}, {offset: 3}}},
		{"abc 123", []Token{{Type: TK_ID, Literal: "abc"}, {Type: TK_INTEGER, Literal: "123"}, {Type: EOF, Literal: ""}}, []Pos{{offset: 0}, {offset: 4}, {offset: 7}}},
		{"(abc 123)", []Token{{Type: TK_LP, Literal: "("}, {Type: TK_ID, Literal: "abc"}, {Type: TK_INTEGER, Literal: "123"}, {Type: TK_RP, Literal: ")"}, {Type: EOF, Literal: ""}}, []Pos{{offset: 0}, {offset: 1}, {offset: 5}, {offset: 8}, {offset: 9}}},
		{"SELECT * FROM abc;", []Token{{Type: TK_SELECT, Literal: "SELECT"}, {Type: TK_STAR, Literal: "*"}, {Type: TK_FROM, Literal: "FROM"}, {Type: TK_ID, Literal: "abc"}, {Type: TK_SEMI, Literal: ";"}, {Type: EOF, Literal: ""}}, []Pos{{offset: 0}, {offset: 7}, {offset: 9}, {offset: 14}, {offset: 17}, {offset: 18}}},
		{"CREATE TABLE abc (id INTEGER);", []Token{{Type: TK_CREATE, Literal: "CREATE"}, {Type: TK_TABLE, Literal: "TABLE"}, {Type: TK_ID, Literal: "abc"}, {Type: TK_LP, Literal: "("}, {Type: TK_ID, Literal: "id"}, {Type: TK_INTEGER, Literal: "INTEGER"}, {Type: TK_RP, Literal: ")"}, {Type: TK_SEMI, Literal: ";"}, {Type: EOF, Literal: ""}}, []Pos{{offset: 0}, {offset: 7}, {offset: 13}, {offset: 17}, {offset: 18}, {offset: 21}, {offset: 28}, {offset: 29}, {offset: 30}}},
	}

	for _, test := range tests {
		tokenizer := New(test.input)

		var tokens []Token
		var positions []Pos

		tok, pos := tokenizer.GetToken()
		tokens = append(tokens, tok)
		positions = append(positions, pos)
		for tokens[len(tokens)-1].Type != EOF {
			tok, pos = tokenizer.GetToken()
			tokens = append(tokens, tok)
			positions = append(positions, pos)
		}

		if !reflect.DeepEqual(tokens, test.expectedIdentifiers) {
			t.Errorf("Expected %v, got %v", test.expectedIdentifiers, tokens)
		}

		if !reflect.DeepEqual(positions, test.positions) {
			t.Errorf("Expected position to be %v, got %v", test.positions, positions)
		}

	}
}
