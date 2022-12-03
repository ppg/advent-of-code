package main

import (
	"testing"

	framework "github.com/ppg/advent-of-code/2022/12/framework"
)

func BenchmarkSolution0(b *testing.B)      { framework.Benchmark(b, parser, solution0, 1) }
func BenchmarkSolution1Part1(b *testing.B) { framework.Benchmark(b, parser, solution1part1, 1) }
func BenchmarkSolution1Part2(b *testing.B) { framework.Benchmark(b, parser, solution1part2, 2) }
