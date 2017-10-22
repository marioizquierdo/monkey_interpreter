package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/marioizquierdo/monkey_interpreter/lib/lexer"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	fmt.Println("Monkey Programming Language")
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		ok := scanner.Scan()
		if !ok {
			return
		}
		line := scanner.Text()
		switch line {
		case "exit":
			return
		case "help":
			fmt.Println("The Monkey Programming Language is designed for the book")
			fmt.Println("'Writing An Interpreter In Go' by Thorsten Ball.")
			fmt.Println("Checkout the book at: https://interpreterbook.com/")
		default:
			l := lexer.New(line)
			for tok := l.NextToken(); tok.Type != lexer.EOF; tok = l.NextToken() {
				fmt.Printf("%v\n", tok)
			}
		}
	}
}
