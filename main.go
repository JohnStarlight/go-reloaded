package main

import (
	"fmt"
	"os"
)

func main() {
	// need exactly two args: input file and output file
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . /path/to/input.txt /path/to/output.txt")
		return
	}
	source, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Usage: go run . /path/to/input.txt /path/to/output.txt")
		return
	}

	// pipeline: tokenize → fix a/an → apply commands → rebuild text
	tokens := Tokenizer(string(source))
	tokens = Vowels(tokens)
	tokens = Transform(tokens)
	result := BuildOutput(tokens)

	err = os.WriteFile(os.Args[2], []byte(result), 0o644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		return
	}
}
