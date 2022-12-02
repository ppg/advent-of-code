package framework

import (
	"bufio"
	"io"
	"testing"
)

func Benchmark[T any](b *testing.B, parse SolutionParser[T], handle SolutionHandler[T], part int) {
	w := io.Discard

	// Read the input
	file, close := readInput(w)
	defer close()

	// Reset the timer and run our loop
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := file.Seek(0, io.SeekStart); err != nil {
			panic(err)
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		s := &solution[T]{handle: handle, parse: parse}
		s.Handle(w, scanner, part)
	}
}
