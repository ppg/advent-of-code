package main

import (
	"io"
	"testing"

	framework "github.com/ppg/advent-of-code/2022/12/framework"
)

func BenchmarkParseFS(b *testing.B) {
	sol := func(w io.Writer, runner *framework.Runner[string]) { parseFS(w, runner) }
	framework.Benchmark(b, parser, sol, 1)
}
func BenchmarkSolution0(b *testing.B) { framework.Benchmark(b, parser, solution0, 1) }
func BenchmarkSolution1(b *testing.B) { framework.Benchmark(b, parser, solution1, 1) }
