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
	framework.Run(os.Stdout)
}

var parser = framework.LineParser

func solution0(w io.Writer, runner *framework.Runner[string]) {
	datastream := <-runner.Lines()
	for i := 4; i < len(datastream); i++ {
		// naive solution to check each 4 char packet
		// TODO: when we check the packet we can jump farther ahead depending on where repeats are;
		// i.e. abccdef after checking abcc we can move to cdef right away
		packet := datastream[i-4 : i]
		fmt.Fprintf(w, "inspecting packet: %s\n", packet)
		m := make(map[rune]bool, 4)
		for _, char := range packet {
			m[char] = true
		}
		if len(m) == 4 {
			fmt.Fprintf(w, "found start-of-packet %s: %d\n", packet, i)
			return
		}
	}
	fmt.Fprintf(w, "could not find start-of-packet\n")
	os.Exit(1)
}
