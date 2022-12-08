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

// 30373
// 25512
// 65332
// 33549
// 35390
var parser = func(line string) []int {
	out := make([]int, len(line))
	for i, height := range line {
		out[i] = int(height - '0')
	}
	return out
}

func solution0(w io.Writer, runner *framework.Runner[[]int]) {
	heights := make(Matrix, 0, 1000)
	for line := range runner.Lines() {
		heights = append(heights, line)
	}
	fmt.Fprintf(w, "heights:\n%s", heights)

	// All the edges are visible to start:
	// (# of rows * 2) + ((# of cols - start - end) * 2)
	visible := 2*len(heights) + 2*(len(heights[0])-2)
	fmt.Fprintf(w, "visible(1): %d\n", visible)

	// Inspect each inner tree to see if it's visible
	// NOTE: naive implementation with inefficient loops
	for row := 1; row < len(heights)-1; row++ {
		for col := 1; col < len(heights[0])-1; col++ {
			// left
			left := true
			for j := col - 1; j >= 0; j-- {
				if heights[row][j] >= heights[row][col] {
					left = false
					break
				}
			}
			// right
			right := true
			for j := col + 1; j < len(heights[row]); j++ {
				if heights[row][j] >= heights[row][col] {
					right = false
					break
				}
			}
			// up
			up := true
			for i := row - 1; i >= 0; i-- {
				if heights[i][col] >= heights[row][col] {
					up = false
					break
				}
			}
			// down
			down := true
			for i := row + 1; i < len(heights); i++ {
				if heights[i][col] >= heights[row][col] {
					down = false
					break
				}
			}
			seen := left || right || up || down
			fmt.Fprintf(w, "%2d: %5t (%2d,%2d) (left=%t,right=%t,up=%t,down=%t)\n",
				heights[row][col], seen, row, col, left, right, up, down)
			if seen {
				visible++
			}
		}
	}
	fmt.Fprintf(w, "visible(2): %d\n", visible)
}

type Matrix [][]int

func (m Matrix) String() string {
	outer := make([]string, len(m))
	for i, row := range m {
		inner := make([]string, len(row))
		for j, col := range row {
			inner[j] = strconv.Itoa(col)
		}
		outer[i] = strings.Join(inner, "") + "\n"
	}
	return strings.Join(outer, "")
}
