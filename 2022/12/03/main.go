package main

import (
	"fmt"
	"io"
	"os"

	framework "github.com/ppg/advent-of-code/2022/12/framework"
)

func main() {
	framework.Register(parser, solution0)
	framework.Run(os.Stdout)
}

var parser = framework.LineParser

func solution0(w io.Writer, runner *framework.Runner[string]) {
	var total int
	fmt.Fprintf(w, "%20s %20s %8s %8s\n", "rucksack1", "rucksack2", "mismatch", "priority")
	for rucksacks := range runner.Lines() {
		rucksack1 := rucksacks[0 : len(rucksacks)/2]
		rucksack2 := rucksacks[len(rucksacks)/2:]
		var mismatch byte
		// O(n^2)
		for i := 0; i < len(rucksack1); i++ {
			for j := 0; j < len(rucksack2); j++ {
				if rucksack1[i] == rucksack2[j] {
					mismatch = rucksack1[i]
					break
				}
			}
		}
		var priority int
		if mismatch > 96 {
			priority = int(mismatch) - 96
		} else {
			priority = int(mismatch) - 38
		}
		fmt.Fprintf(w, "%20s %20s %8s %8d\n", rucksack1, rucksack2, fmt.Sprintf("%s (%3d)", string([]byte{mismatch}), mismatch), priority)

		total += priority
	}
	fmt.Fprintf(w, "Total: %d\n", total)
}
