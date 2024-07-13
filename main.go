package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/rqlite/sql"
)

const (
	PROMPT = ">> "
)

func main() {
	// Start REPL
	Start(os.Stdin, os.Stdout)
}

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

		parser := sql.NewParser(strings.NewReader(line))
		stmt, err := parser.ParseStatement()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(stmt.String())
	}
}
