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

		if line == "exit" {
			fmt.Println("Goodbye!")
			return
		}

		t := tokenizer.New(line)

		for tok, _ := t.GetToken(); tok.Type != tokenizer.EOF; tok, _ = t.GetToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
