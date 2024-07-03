package parser

import (
	"testing"

	"kittendb/tokenizer"
)

func TestSelectStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"SELECT * FROM abc", "SELECT * FROM abc"},
	}

	for _, test := range tests {
		tokenizer := tokenizer.New(test.input)
		p := New(tokenizer)

		stmt := p.EvalSQL(test.input)
		checkParserErrors(t, p)
		if test.expected != stmt.String() {
			t.Fatalf("expected=%q, got=%q", test.expected, stmt.String())
		}
	}
}
