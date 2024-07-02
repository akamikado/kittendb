package repl

import (
	"bufio"
	"fmt"
	"io"

	"kittendb/tokenizer"
)

const (
	PROMPT = ">> "
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		t := tokenizer.New(line)

		for tok := t.GetToken(); tok.Type != tokenizer.EOF; tok = t.GetToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
