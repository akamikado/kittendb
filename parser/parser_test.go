package parser

import (
	"fmt"
	"testing"
)

func TestSQLStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"SELECT * FROM abc", "SELECT * FROM abc"},
		{"SELECT name FROM patients", "SELECT name FROM patients"},
		{"SELECT name, age FROM patients", "SELECT name, age FROM patients"},
		{"SELECT DISTINCT name FROM patients", "SELECT DISTINCT name FROM patients"},
		{"INSERT INTO users (name, age) VALUES ('John', 30)", "INSERT INTO users (name, age) VALUES ('John', 30)"},
		{"INSERT INTO users VALUES ('John', 30)", "INSERT INTO users VALUES ('John', 30)"},
	}

	for _, test := range tests {
		p := New(test.input)

		stmt, err := p.ParseStatement()
		if err != nil {
		}
		if test.expected != stmt.String() {
			fmt.Println("got ", stmt)
			t.Errorf("expected=%q, got=%q", test.expected, stmt.String())
		}
	}
}
