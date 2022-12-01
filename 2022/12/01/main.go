/*
https://adventofcode.com/2022/day/1
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	elves := make([]*Elf, 0, 256) // 256 is the input.txt size
	elf := new(Elf)
	elves = append(elves, elf)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())
		if line == "" {
			fmt.Printf("%d: %d total\n", elf.id, elf.calories)
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

	fmt.Printf("Highest Calories Elves\n")
	var total int
	for i := 0; i < 3; i++ {
		elf := elves[len(elves)-1-i]
		fmt.Printf("Elf %d: %d\n", elf.id, elf.calories)
		total += elf.calories
	}
	fmt.Printf("Total: %d\n", total)

}

type Elf struct {
	id       int
	calories int
}
