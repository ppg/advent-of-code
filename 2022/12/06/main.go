/*
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
	framework.Register(parser, solution1)
	framework.Run(os.Stdout)
}

var parser = framework.LineParser

func solution0(w io.Writer, runner *framework.Runner[string]) {
	datastream := <-runner.Lines()
	var markerSize int
	runner.ByPart(
		func() { markerSize = 4 },
		func() { markerSize = 14 },
	)
	for i := markerSize; i < len(datastream); i++ {
		packet := datastream[i-markerSize : i]
		fmt.Fprintf(w, "inspecting packet: %s\n", packet)

		// naive solution to check each markerSize char packet
		// TODO: when we check the packet we can jump farther ahead depending on where repeats are;
		// i.e. abccdef after checking abcc we can move to cdef right away
		m := make(map[rune]bool, markerSize)
		for _, char := range packet {
			m[char] = true
		}
		if len(m) == markerSize {
			fmt.Fprintf(w, "found start-of-packet %s: %d\n", packet, i)
			return
		}
	}
	fmt.Fprintf(w, "could not find start-of-packet\n")
	os.Exit(1)
}

func solution1(w io.Writer, runner *framework.Runner[string]) {
	datastream := <-runner.Lines()
	var markerSize int
	runner.ByPart(
		func() { markerSize = 4 },
		func() { markerSize = 14 },
	)
	i := markerSize
	for i <= len(datastream) {
		packet := datastream[i-markerSize : i]
		fmt.Fprintf(w, "inspecting packet: %s\n", packet)

		// Search for the latest repeated char and advance to that position
		// i.e. abccdef after checking abcc we can move to cdef right away
		var dupes bool
		m := make(map[byte]bool, markerSize)
		for j := len(packet) - 1; j >= 0; j-- {
			char := packet[j]
			if m[char] {
				// advance 1 past where we found the matching char since we don't need to retest that
				i = i + j + 1
				dupes = true
				break
			}
			m[char] = true
		}
		// If no dupes, we found it
		if !dupes {
			fmt.Fprintf(w, "found start-of-packet %s: %d\n", packet, i)
			return
		}
	}
	fmt.Fprintf(w, "could not find start-of-packet\n")
	os.Exit(1)
}
