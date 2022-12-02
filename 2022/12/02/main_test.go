package main

import (
	"testing"

	framework "github.com/ppg/advent-of-code/2022/12/framework"
)

func BenchmarkSolution0Part1(b *testing.B) { framework.Benchmark(b, parser, solution0, 1) }
func BenchmarkSolution0Part2(b *testing.B) { framework.Benchmark(b, parser, solution0, 2) }
func BenchmarkSolution1Part1(b *testing.B) { framework.Benchmark(b, parser, solution1, 1) }
func BenchmarkSolution2Part2(b *testing.B) { framework.Benchmark(b, parser, solution1, 2) }
