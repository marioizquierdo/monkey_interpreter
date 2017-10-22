package main

import (
	"os"

	"github.com/marioizquierdo/monkey_interpreter/lib/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
