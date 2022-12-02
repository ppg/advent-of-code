package framework

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
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

func LineParser(line string) string { return strings.TrimSpace(line) }

func ColParser(n int) func(line string) []string {
	return func(line string) []string {
		cols := strings.Split(line, " ")
		if len(cols) != n {
			panic(fmt.Errorf("unexpected number of columns %d, expected %d:\n%s", len(cols), n, line))
		}
		return cols
	}
}
