package main

import (
	"testing"

	framework "github.com/ppg/advent-of-code/2022/12/framework"
)

func BenchmarkSolution0(b *testing.B) { framework.Benchmark(b, solution0, 2) }
func BenchmarkSolution1(b *testing.B) { framework.Benchmark(b, solution1, 2) }
