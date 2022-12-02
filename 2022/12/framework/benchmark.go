package framework

import (
	"bufio"
	"io"
	"testing"
)

func Benchmark(b *testing.B, solution Solution, part int) {
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
		solution(w, &Runner{scanner: scanner, part: part})
	}
}
