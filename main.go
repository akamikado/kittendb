package main

import (
	"os"

	"kittendb/repl"
)

func main() {
	// Start REPL
	repl.Start(os.Stdin, os.Stdout)
}
