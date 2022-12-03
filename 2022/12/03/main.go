/*
Part 1:
Total: 8072
Part 2:
Total: 2567
*/
package main

import (
	"fmt"
	"io"
	"os"

	framework "github.com/ppg/advent-of-code/2022/12/framework"
)

func main() {
	framework.Register(parser, solution0)
	framework.Register(parser, solution1part1)
	framework.Register(parser, solution1part2)
	framework.Run(os.Stdout)
}

var parser = framework.LineParser

// solution0 solves part 1 in O(n^2) time with a nested search on each rucksack
func solution0(w io.Writer, runner *framework.Runner[string]) {
	var total int
	fmt.Fprintf(w, "%20s %20s %8s %8s\n", "rucksack1", "rucksack2", "mismatch", "priority")
	for rucksacks := range runner.Lines() {
		rucksack1 := rucksacks[0 : len(rucksacks)/2]
		rucksack2 := rucksacks[len(rucksacks)/2:]

		// O(n^2)
		var mismatch rune
		for i := 0; i < len(rucksack1); i++ {
			for j := 0; j < len(rucksack2); j++ {
				if rucksack1[i] == rucksack2[j] {
					mismatch = rune(rucksack1[i])
					break
				}
			}
		}

		pri := priority(mismatch)
		fmt.Fprintf(w, "%20s %20s %8s %8d\n", rucksack1, rucksack2, fmt.Sprintf("%s (%3d)", string([]rune{mismatch}), mismatch), pri)

		total += pri
	}
	fmt.Fprintf(w, "Total: %d\n", total)
}

// solution1part1 solves part 1 in O(log n) + O(n) time by converting the
// rucksack to a hash (O(log n)) and then searching each item in the first
// rucksack (O(n)) to find if the second rucksack hash has it (O(1)).
func solution1part1(w io.Writer, runner *framework.Runner[string]) {
	var total int
	fmt.Fprintf(w, "%20s %20s %8s %8s\n", "rucksack1", "rucksack2", "mismatch", "priority")
	for rucksacks := range runner.Lines() {
		rucksack1 := rucksacks[0 : len(rucksacks)/2]
		rucksack2 := rucksacks[len(rucksacks)/2:]

		mr1 := make(map[rune]bool, len(rucksack1))
		for _, item := range rucksack1 {
			mr1[item] = true
		}
		mr2 := make(map[rune]bool, len(rucksack2))
		for _, item := range rucksack2 {
			mr2[item] = true
		}

		// O(?)
		var mismatch rune
		for item := range mr1 {
			if mr2[item] {
				mismatch = item
				break
			}
		}

		pri := priority(mismatch)
		fmt.Fprintf(w, "%20s %20s %8s %8d\n", rucksack1, rucksack2, fmt.Sprintf("%s (%3d)", string([]rune{mismatch}), mismatch), pri)

		total += pri
	}
	fmt.Fprintf(w, "Total: %d\n", total)
}

// solution1part2 solves part 2 in O(log n) + O(n) time by converting the
// group'srucksacks to a hash (O(log n)) and then searching each item in the first
// rucksack (O(n)) to find if the second and third rucksack hashes have it (O(1)).
func solution1part2(w io.Writer, runner *framework.Runner[string]) {
	var total int
	rucksacks := make([]string, 0, 1000)
	rucksackHashes := make([]map[rune]bool, 0, 1000)
	for rucksack := range runner.Lines() {
		mr := make(map[rune]bool, len(rucksack))
		for _, item := range rucksack {
			mr[item] = true
		}
		rucksacks = append(rucksacks, rucksack)
		rucksackHashes = append(rucksackHashes, mr)
	}

	fmt.Fprintf(w, "%40s %40s %40s %8s %8s\n",
		"rucksack1", "rucksack2", "rucksack3", "badge", "priority")
	for i := 0; i < len(rucksackHashes); i += 3 {
		// O(?)
		var badge rune
		for item := range rucksackHashes[i] {
			if !rucksackHashes[i+1][item] {
				continue
			}
			if !rucksackHashes[i+2][item] {
				continue
			}
			badge = item
		}

		pri := priority(badge)
		fmt.Fprintf(w, "%40s %40s %40s %8s %8d\n",
			rucksacks[i],
			rucksacks[i+1],
			rucksacks[i+2],
			fmt.Sprintf("%s (%3d)", string([]rune{badge}), badge),
			pri)

		total += pri
	}
	fmt.Fprintf(w, "Total: %d\n", total)
}

func priority(item rune) int {
	if item > 96 {
		return int(item) - 96
	}
	return int(item) - 38
}
