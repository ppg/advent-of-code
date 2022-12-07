/*
https://adventofcode.com/2022/day/1

	Highest Calories Elves
	Elf 151: 74394
	Elf 191: 69863
	Elf 99: 68579
	Total: 212836
*/
package main

import (
	"container/heap"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"

	"github.com/ppg/advent-of-code/2022/12/framework"
)

func main() {
	framework.Register(framework.LineParser, solution0)
	framework.Register(framework.LineParser, solution1)
	framework.Run(os.Stdout)
}

// Solution 0 collects the elves into a slice and sorts at the end.
// This solution handles both parts of the question.
func solution0(w io.Writer, runner *framework.Runner[string]) {
	elves := make([]*Elf, 0, 256) // 256 is the input.txt size
	elf := new(Elf)
	elves = append(elves, elf)
	for line := range runner.Lines() {
		if line == "" {
			fmt.Fprintf(w, "%d: %d total\n", elf.id, elf.calories)
			elf = &Elf{id: elf.id + 1}
			elves = append(elves, elf)
		} else {
			calories, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			elf.calories += calories
		}
	}

	sort.SliceStable(elves, func(i, j int) bool { return elves[i].calories < elves[j].calories })

	fmt.Fprintf(w, "Highest Calories Elves (count:%d)\n", len(elves))
	var total int
	for i := 0; i < 3; i++ {
		elf := elves[len(elves)-1-i]
		fmt.Fprintf(w, "Elf %d: %d\n", elf.id, elf.calories)
		total += elf.calories
	}
	fmt.Fprintf(w, "Total: %d\n", total)
}

type Elf struct {
	id       int
	calories int
}

// Less prioritizes more calories so the top of the heap is the largest calories
func (e Elf) Less(other Elf) bool { return e.calories > other.calories }

// Solution 1 uses a heap to track the highest elves in an ongoing, efficient fashion.
// This solution handles both parts of the question.
func solution1(w io.Writer, runner *framework.Runner[string]) {
	elves := make(framework.Heap[Elf], 0, 256) // 256 is the input.txt size
	var elf Elf
	for line := range runner.Lines() {
		if line == "" {
			// Add elf to the heap
			fmt.Fprintf(w, "%d: %d total\n", elf.id, elf.calories)
			heap.Push(&elves, elf)

			// Create new elft
			elf = Elf{id: elf.id + 1}
		} else {
			// Accumulate elf calories
			calories, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			elf.calories += calories
		}
	}

	fmt.Fprintf(w, "Highest Calories Elves (count:%d)\n", len(elves))
	var total int
	for i := 0; i < 3; i++ {
		elf := heap.Pop(&elves).(Elf)
		fmt.Fprintf(w, "Elf %03d: %d\n", elf.id, elf.calories)
		total += elf.calories
	}
	fmt.Fprintf(w, "Total: %d\n", total)
}
