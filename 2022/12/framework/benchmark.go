package framework

import (
	"bufio"
	"io"
	"testing"
)

func Benchmark[T any](b *testing.B, parse SolutionParser[T], handle SolutionHandler[T], part int) {
	w := io.Discard

	// Read the test input
	reader := readTestInput(w)
	defer reader.Close()

	// Reset the timer and run our loop
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := reader.Seek(0, io.SeekStart); err != nil {
			panic(err)
		}
		scanner := bufio.NewScanner(reader)
		scanner.Split(bufio.ScanLines)
		s := &solution[T]{handle: handle, parse: parse}
		s.Handle(w, scanner, part)
	}
}
