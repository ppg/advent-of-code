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
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	h := make(ElfHeap, 0, 256) // 256 is the input.txt size
	var elf Elf

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())
		if line == "" {
			// Add elf to the heap
			if os.Getenv("DEBUG") != "" {
				fmt.Printf("%d: %d total\n", elf.id, elf.calories)
			}
			heap.Push(&h, elf)

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

	fmt.Printf("Highest Calories Elves\n")
	var total int
	for i := 0; i < 3; i++ {
		elf := heap.Pop(&h).(Elf)
		fmt.Printf("Elf %03d: %d\n", elf.id, elf.calories)
		total += elf.calories
	}
	fmt.Printf("Total: %d\n", total)
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

type Elf struct {
	id       int
	calories int
}
