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
	"strings"

	"github.com/ppg/advent-of-code/2022/12/framework"
)

func main() {
	framework.Register(solution0)
	framework.Register(solution1)
	framework.Run(os.Stdout)
}

// Solution 0 collects the elves into a slice and sorts at the end.
// This solution handles both parts of the question.
func solution0(w io.Writer, runner *framework.Runner) {
	elves := make([]*Elf, 0, 256) // 256 is the input.txt size
	elf := new(Elf)
	elves = append(elves, elf)
	for runner.Scan() {
		line := strings.TrimSpace(runner.Text())
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

// Solution 1 uses a heap to track the highest elves in an ongoing, efficient fashion.
// This solution handles both parts of the question.
func solution1(w io.Writer, runner *framework.Runner) {
	elves := make(ElfHeap, 0, 256) // 256 is the input.txt size
	var elf Elf
	for runner.Scan() {
		line := strings.TrimSpace(runner.Text())
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

type ElfHeap []Elf

func (h ElfHeap) Len() int { return len(h) }

// Less prioritizes more calories so the top of the heap is the largest calories
func (h ElfHeap) Less(i, j int) bool { return h[i].calories > h[j].calories }
func (h ElfHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ElfHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Elf))
}

func (h *ElfHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
