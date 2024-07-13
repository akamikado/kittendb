package parser

import "fmt"

type Pos struct {
	offset int
	line   int
	column int
}

func (p Pos) String() string {
	return fmt.Sprintf("%d", p.offset)
}

func (p Pos) IsValid() bool {
	return p != Pos{}
}

type Tokenizer struct {
	input        []byte
	readPosition Pos
	char         byte
}

func NewTokenizer(input string) *Tokenizer {
	inputInBytes := []byte(input)
	t := &Tokenizer{input: inputInBytes}
	t.char = t.input[t.readPosition.offset]
	return t
}

// Todo: Fix implementation of numbers
// Todo: Fix whitespace and new line handling
func (t *Tokenizer) GetToken() (Token, Pos) {
	var token Token
	var pos Pos

	if t.char == ' ' {
		for ; t.readPosition.offset < len(t.input) && t.input[t.readPosition.offset] == ' '; t.readPosition.offset += 1 {
			t.readPosition.column += 1
		}
		if t.readPosition.offset < len(t.input) {
			t.char = t.input[t.readPosition.offset]
		} else {
			t.char = 0
		}
	} else if t.char == '\n' {
		for ; t.readPosition.offset < len(t.input) && t.input[t.readPosition.offset] == '\n'; t.readPosition.offset += 1 {
			t.readPosition.line += 1
			t.readPosition.column = 1
		}
		if t.readPosition.offset < len(t.input) {
			t.char = t.input[t.readPosition.offset]
		} else {
			t.char = 0
		}
	}

	switch t.char {
	case 0:
		token.Type = EOF
		token.Literal = ""
		pos.offset = t.readPosition.offset
		return token, pos
	case ' ':
		t.readPosition.offset += 1
		t.readPosition.column += 1
		t.char = t.input[t.readPosition.offset]
		token, pos = t.GetToken()
	case '\n':
		t.readPosition.offset += 1
		t.readPosition.column = 1
		t.readPosition.line += 1
		t.char = t.input[t.readPosition.offset]
		token, pos = t.GetToken()
	case ';':
		token.Type = TK_SEMI
		token.Literal = ";"
		pos.offset = t.readPosition.offset
	case '(':
		token.Type = TK_LP
		token.Literal = "("
		pos.offset = t.readPosition.offset
	case ')':
		token.Type = TK_RP
		token.Literal = ")"
		pos.offset = t.readPosition.offset
	case '*':
		token.Type = TK_STAR
		token.Literal = "*"
		pos.offset = t.readPosition.offset
	case ',':
		token.Type = TK_COMMA
		token.Literal = ","
		pos.offset = t.readPosition.offset
	default:
		if isAlphabet(t.char) {
			pos.offset = t.readPosition.offset
			token.Literal = t.ReadIdentifier()
			token.Type = LookupKeyword(token.Literal)
		} else if isDigit(t.char) {
			pos.offset = t.readPosition.offset
			token.Literal = t.ReadNumber()
			token.Type = TK_INTEGER
		} else {
			pos.offset = t.readPosition.offset
			token.Type = TK_ILLEGAL
		}
	}

	t.readPosition.offset += 1
	t.readPosition.column += 1
	if t.readPosition.offset < len(t.input) {
		t.char = t.input[t.readPosition.offset]
	} else {
		t.char = 0
	}

	return token, pos
}

func (t *Tokenizer) PeekByte() byte {
	if t.readPosition.offset >= len(t.input)-1 {
		return 0
	}

	return t.input[t.readPosition.offset+1]
}

func isAlphabet(c byte) bool {
	if ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') {
		return true
	}
	return false
}

func isDigit(c byte) bool {
	if '0' <= c && c <= '9' {
		return true
	}
	return false
}

func (t *Tokenizer) ReadNumber() string {
	position := t.readPosition.offset
	for ; t.readPosition.offset < len(t.input) && isDigit(t.input[t.readPosition.offset]); t.readPosition.offset += 1 {
	}
	t.readPosition.offset -= 1
	return string(t.input[position : t.readPosition.offset+1])
}

func (t *Tokenizer) ReadIdentifier() string {
	position := t.readPosition.offset
	for ; t.readPosition.offset < len(t.input) && isAlphabet(t.input[t.readPosition.offset]); t.readPosition.offset += 1 {
	}
	t.readPosition.offset -= 1
	return string(t.input[position : t.readPosition.offset+1])
}

func (t *Tokenizer) PeekToken() Token {
	pos := t.readPosition
	char := t.char
	token, _ := t.GetToken()
	t.char = char
	t.readPosition = pos
	return token
}

func (t *Tokenizer) GetPos() Pos {
	return t.readPosition
}
