/*
 */
package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	framework "github.com/ppg/advent-of-code/2022/12/framework"
)

func main() {
	framework.Register(parser, solution0)
	framework.Run(os.Stdout)
}

var parser = framework.LineParser

func solution0(w io.Writer, runner *framework.Runner[string]) {
	// Parse insructions
	// noop
	// addx 3
	// addx -5

	var (
		cycle int
		x     int = 1
		// []int => cycle, reg
		samples          []int
		sampleSignals    []int
		sampleSignalsSum int
	)
	positions := make(map[int]int)
	for line := range runner.Lines() {
		// ** before cycle

		// ** during cycle
		var value int
		// Parse instruction
		matches := strings.Split(line, " ")
		switch matches[0] {
		case "noop":
			cycle++
		case "addx":
			value = parseInt(matches[1])
			cycle += 2
		default:
			panic(fmt.Errorf("unrecognized command: %s", line))
		}

		// start sampling every 40 cycles after cycle 20
		if cycle >= 20 {
			sampleIndex := (cycle - 20) / 40
			sampleCycle := sampleIndex*40 + 20
			for i := len(samples); i <= sampleIndex; i++ {
				sampleSignal := x * sampleCycle
				fmt.Fprintf(w, "sampling %d: %d => %d\n", sampleCycle, x, sampleSignal)
				samples = append(samples, x)
				sampleSignals = append(sampleSignals, sampleSignal)
				sampleSignalsSum += sampleSignal
			}
		}

		// ** after cycle
		// increment register after the cycle has completed
		x += value
		fmt.Fprintf(w, "%4d: %3d - %s\n", cycle, x, line)
		positions[cycle] = x
	}
	fmt.Fprintf(w, "sample signal strength sum: %d\n", sampleSignalsSum)

	// draw
	// sprite starts at 1
	// 0123456789012345678901234567890123456789
	// ###.....................................
	pos := 1
	for row := 0; row < 6; row++ {
		for i := 0; i < 40; i++ {
			cycle := row*40 + i + 1

			// 0123456789012345678901234567890123456789
			// ...............###......................
			pixel := "."
			if i >= pos-1 && i <= pos+1 {
				pixel = "#"
			}
			fmt.Fprint(w, pixel)

			// If we have a new sprite position adjust
			if newPos, ok := positions[cycle]; ok {
				pos = newPos
			}
		}
		fmt.Fprintln(w)
	}
}

// TODO(ppg): framework candidate
func parseInt(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return out
}
