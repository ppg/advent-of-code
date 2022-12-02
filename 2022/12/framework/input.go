package framework

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func readInput(w io.Writer) (*os.File, func() error) {
	flag.Parse()
	input := flag.Arg(0)
	if input == "" {
		input = "input.txt"
	}
	fmt.Fprintf(w, "reading input: %s\n", input)
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	return file, file.Close
}
