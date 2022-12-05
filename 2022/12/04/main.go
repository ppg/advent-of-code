/*
 */
package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	framework "github.com/ppg/advent-of-code/2022/12/framework"
)

func main() {
	framework.Register(parser, solution0)
	framework.Run(os.Stdout)
}

var rePairs = regexp.MustCompile("([0-9]+)-([0-9]+),([0-9]+)-([0-9]+)")

var parser = func(line string) [2]section {
	matches := rePairs.FindStringSubmatch(line)
	return [2]section{
		{start: parseInt(matches[1]), end: parseInt(matches[2])},
		{start: parseInt(matches[3]), end: parseInt(matches[4])},
	}
}

func solution0(w io.Writer, runner *framework.Runner[[2]section]) {
	// Read in cleaning pairs
	//   2-4,6-8
	//   2-3,4-5
	//   5-7,7-9
	//   2-8,3-7
	//   6-6,4-6
	//   2-6,4-8
	var total, overlapping int
	for sections := range runner.Lines() {
		elf1 := sections[0]
		elf2 := sections[1]
		fmt.Fprintf(w, "elf1: %s\n", elf1.Format(max(elf1.end, elf2.end)))
		fmt.Fprintf(w, "elf2: %s", elf2.Format(max(elf1.end, elf2.end)))
		var overlap bool
		runner.ByPart(
			func() { overlap = elf1.overlap(elf2) || elf2.overlap(elf1) },
			func() {
				//   ..234....
				//   ......678
				//   ..23..
				//   ....45
				//
				//   .....567..
				//   .......789
				overlap = overlap || elf1.start <= elf2.start && elf1.end >= elf2.start
				//   ......6
				//   ....456
				overlap = overlap || elf1.end >= elf2.end && elf1.start <= elf2.end

				// flip of the first 2
				overlap = overlap || elf2.start <= elf1.start && elf2.end >= elf1.start
				overlap = overlap || elf2.end >= elf1.end && elf2.start <= elf1.end
			},
		)
		total++
		if overlap {
			fmt.Fprintf(w, " overlap\n")
			overlapping++
		} else {
			fmt.Fprintf(w, " non-overlap\n")
		}
	}
	fmt.Fprintf(w, "Total: %d\n", total)
	fmt.Fprintf(w, "Overlapping: %d\n", overlapping)
}

type section struct {
	start, end int
}

func (s section) overlap(other section) bool {
	return s.start <= other.start && s.end >= other.end
}

func (s section) String() string { return fmt.Sprintf("%d-%d", s.start, s.end) }

func (s section) Format(max int) string {
	before := strings.Repeat(".", s.start)
	during := ""
	for i := s.start; i <= s.end; i++ {
		during += fmt.Sprintf("%d", i)
	}
	after := strings.Repeat(".", max-s.end)
	return before + during + after
}

// TODO(ppg): move to framework
func parseInt(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return out
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
