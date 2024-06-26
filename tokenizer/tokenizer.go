package tokenizer

type Tokenizer struct {
	input        []byte
	readPosition int
	char         byte
}

func New(input []byte) *Tokenizer {
	t := &Tokenizer{input: input}
	t.char = t.input[t.readPosition]
	return t
}

// Todo: Fix implementation of numbers
// Todo: Fix whitespace and new line handling
func (t *Tokenizer) GetToken() Token {
	var token Token

	if t.char == ' ' {
		for ; t.readPosition < len(t.input) && t.input[t.readPosition] == ' '; t.readPosition += 1 {
		}
		if t.readPosition < len(t.input) {
			t.char = t.input[t.readPosition]
		} else {
			t.char = 0
		}
	} else if t.char == '\n' {
		for ; t.readPosition < len(t.input) && t.input[t.readPosition] == '\n'; t.readPosition += 1 {
		}
		if t.readPosition < len(t.input) {
			t.char = t.input[t.readPosition]
		} else {
			t.char = 0
		}
	}

	switch t.char {
	case 0:
		token.Type = EOF
		token.Literal = ""
		return token
	case ' ':
		t.readPosition += 1
		t.char = t.input[t.readPosition]
		token = t.GetToken()
	case '\n':
		t.readPosition += 1
		t.char = t.input[t.readPosition]
		token = t.GetToken()
	case ';':
		token.Type = TK_SEMI
		token.Literal = ";"
	case '(':
		token.Type = TK_LP
		token.Literal = "("
	case ')':
		token.Type = TK_RP
		token.Literal = ")"
	case '*':
		token.Type = TK_STAR
		token.Literal = "*"
	default:
		if isAlphabet(t.char) {
			token.Literal = t.ReadIdentifier()
			token.Type = LookupKeyword(token.Literal)
		} else if isDigit(t.char) {
			token.Literal = t.ReadNumber()
			token.Type = TK_INTEGER
		} else {
			token.Type = TK_ILLEGAL
		}
	}

	t.readPosition += 1
	if t.readPosition < len(t.input) {
		t.char = t.input[t.readPosition]
	} else {
		t.char = 0
	}

	return token
}

func (t *Tokenizer) PeekByte() byte {
	if t.readPosition >= len(t.input)-1 {
		return 0
	}

	return t.input[t.readPosition+1]
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
	position := t.readPosition
	for ; t.readPosition < len(t.input) && isDigit(t.input[t.readPosition]); t.readPosition += 1 {
	}
	t.readPosition -= 1
	return string(t.input[position : t.readPosition+1])
}

func (t *Tokenizer) ReadIdentifier() string {
	position := t.readPosition
	for ; t.readPosition < len(t.input) && isAlphabet(t.input[t.readPosition]); t.readPosition += 1 {
	}
	t.readPosition -= 1
	return string(t.input[position : t.readPosition+1])
}
