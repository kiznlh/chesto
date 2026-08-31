package main

import (
	"bufio"
	"fmt"
	"os"
)

// Front end: Tokenizer -> Parser -> Code generator. Input = sql query, Output: sqlite virtual machine bytecode
// Backend: -> Virtual Machine -> B-Tree -> Pager -> OS interface. Input qlite virtual machine bytecode

// Make a REPL

func print_prompt() {
	fmt.Printf("db > ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		print_prompt()

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()

		// Handle input
		if input == ".exit" {
			break
		}

		fmt.Println(input)
	}
}
