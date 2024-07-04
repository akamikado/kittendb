package parser

import (
	"testing"
)

func TestSelectStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"SELECT * FROM abc", "SELECT * FROM abc"},
	}

	for _, test := range tests {
		p := New(test.input)

		stmt := p.ParseStatement()
		checkParserErrors(t, p)
		if test.expected != stmt.String() {
			t.Fatalf("expected=%q, got=%q", test.expected, stmt.String())
		}
	}
}
