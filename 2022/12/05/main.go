/*
 */
package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"

	framework "github.com/ppg/advent-of-code/2022/12/framework"
)

func main() {
	framework.Register(parser, solution0)
	framework.Run(os.Stdout)
}

var parser = func(line string) string { return line }

var reMove = regexp.MustCompile("move ([0-9]+) from ([0-9]+) to ([0-9]+)")

func solution0(w io.Writer, runner *framework.Runner[string]) {
	lines := runner.Lines()
	fmt.Fprintf(w, "Reading stacks...\n")
	// read in stacks
	//       [D]
	//   [N] [C]
	//   [Z] [M] [P]
	var stacks []stack
	for line := range lines {
		//fmt.Printf("line: %s\n", line)
		if line[1] == byte('1') {
			break
		}

		for j := 1; j < len(line); j += 4 {
			col := j / 4
			for i := len(stacks); i <= col; i++ {
				stacks = append(stacks, make(stack, 0, 100))
			}
			if line[j] != byte(' ') {
				stacks[col] = append(stacks[col], string(line[j]))
			}
		}
	}
	//printStacks(w, stacks)
	// reverse so the 'top' of the stack is at the end of the slice
	for _, stack := range stacks {
		for j := 0; j < len(stack)/2; j++ {
			//fmt.Printf("swapping %s and %s\n", stack[j], stack[len(stack)-1-j])
			stack[j], stack[len(stack)-1-j] = stack[len(stack)-1-j], stack[j]
		}
	}
	printStacks(w, stacks)

	// Read empty line
	<-lines
	fmt.Fprintf(w, "\n\n")

	fmt.Fprintf(w, "Performing moves...\n")
	// read in moves
	//   move 1 from 2 to 1
	//   move 3 from 1 to 3
	//   move 2 from 2 to 1
	//   move 1 from 1 to 2
	for line := range lines {
		//fmt.Printf("line: %s\n", line)
		matches := reMove.FindStringSubmatch(line)
		count, err := strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}
		from, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}
		to, err := strconv.Atoi(matches[3])
		if err != nil {
			panic(err)
		}
		// perform operations
		fmt.Fprintf(w, "moving %d from %d to %d\n", count, from, to)
		runner.ByPart(
			func() { // part 1 count > 1 does one at a time (reverses order on the to)
				for i := 0; i < count; i++ {
					stacks[to-1].push(stacks[from-1].pop())
				}
			},
			func() { // part 2 count > 1 does all at once (keeps order on the to)
				var toMove stack
				for i := 0; i < count; i++ {
					toMove.push(stacks[from-1].pop())
				}
				fmt.Fprintf(w, "toMove: %#v\n", toMove)
				for len(toMove) > 0 {
					stacks[to-1].push(toMove.pop())
				}
			},
		)
		printStacks(w, stacks)
	}
	fmt.Fprintf(w, "Final stacks:\n")
	printStacks(w, stacks)

	fmt.Fprintf(w, "First items: ")
	for _, stack := range stacks {
		fmt.Fprintf(w, "%s", stack[len(stack)-1])
	}
	fmt.Fprintf(w, "\n")

}

type stack []string

func (s *stack) push(item string) {
	*s = append(*s, item)
}
func (s *stack) pop() string {
	element := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return element
}

func printStacks(w io.Writer, stacks []stack) {
	for i, stack := range stacks {
		fmt.Fprintf(w, "Stack %02d:", i+1)
		for _, item := range stack {
			fmt.Fprintf(w, " %s", item)
		}
		fmt.Fprintf(w, "\n")
	}
}
