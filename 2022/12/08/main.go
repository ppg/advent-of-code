/*
 */
package main

import (
	"fmt"
	"io"
	"os"
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
	trees := make(Matrix, 0, 1000)
	var (
		rows int
		cols int
	)
	for line := range runner.Lines() {
		rowTrees := make([]*Tree, len(line))
		for col, height := range line {
			rowTrees[col] = &Tree{height: height, row: rows, col: col}
		}
		cols = len(rowTrees)
		trees = append(trees, rowTrees)
		rows++
	}
	//fmt.Fprintf(w, "heights:\n")
	//trees.Format(w, func(t *Tree) string { return strconv.Itoa(t.height) })

	runner.ByPart(
		func() { // part 1
			// All the edges are visible to start:
			// (# of rows * 2) + ((# of cols - start - end) * 2)
			visible := 2*len(trees) + 2*(len(trees[0])-2)
			fmt.Fprintf(w, "visible(1): %d\n", visible)

			// Inspect each inner tree to see if it's visible
			// NOTE: naive implementation with inefficient loops
			for row := 1; row < len(trees)-1; row++ {
				for col := 1; col < len(trees[0])-1; col++ {
					tree := trees[row][col]

					// left
					left := true
					for j := col - 1; j >= 0; j-- {
						if trees[row][j].height >= tree.height {
							left = false
							break
						}
					}
					// right
					right := true
					for j := col + 1; j < len(trees[row]); j++ {
						if trees[row][j].height >= tree.height {
							right = false
							break
						}
					}
					// up
					up := true
					for i := row - 1; i >= 0; i-- {
						if trees[i][col].height >= tree.height {
							up = false
							break
						}
					}
					// down
					down := true
					for i := row + 1; i < len(trees); i++ {
						if trees[i][col].height >= tree.height {
							down = false
							break
						}
					}
					seen := left || right || up || down
					fmt.Fprintf(w, "%2d: %5t (%s) (left=%t,right=%t,up=%t,down=%t)\n",
						tree.height, seen, tree.pos(), left, right, up, down)
					if seen {
						visible++
					}
				}
			}
			fmt.Fprintf(w, "visible(2): %d\n", visible)
		},
		func() { // part 2
			// NOTE: we must mark our neighbors in a given direction before we mark
			// ourselves, so we iterate over the matrix in the four directions.

			// lefts
			//fmt.Fprintf(w, "lefts\n")
			for row := 0; row < rows; row++ {
				for col := 1; col < cols; col++ {
					tree := trees[row][col]
					for j := col - 1; j >= 0; j-- {
						tree.left++
						if trees[row][j].height >= tree.height {
							break
						}
					}
					//neighbor := trees[row][col-1]
					////fmt.Fprintf(w, "\n")
					////fmt.Fprintf(w, "neighbor: (%d,%d) %d - %d\n", neighbor.row, neighbor.col, neighbor.height, neighbor.left)
					//if tree.height > neighbor.height {
					//	tree.left = neighbor.left + 1
					//} else {
					//	tree.left = 1
					//}
					//fmt.Fprintf(w, " left: (%d,%d) %d - %d\n", tree.row, tree.col, tree.height, tree.left)
				}
			}
			// rights
			//fmt.Fprintf(w, "rights\n")
			for row := 0; row < rows; row++ {
				for col := cols - 2; col >= 0; col-- {
					tree := trees[row][col]
					for j := col + 1; j < cols; j++ {
						tree.right++
						if trees[row][j].height >= tree.height {
							break
						}
					}
					//neighbor := trees[row][col+1]
					////fmt.Fprintf(w, "\n")
					////fmt.Fprintf(w, "neighbor: (%s) %d - %d\n", neighbor.pos(), neighbor.height, neighbor.right)
					//if tree.height > neighbor.height {
					//	tree.right = neighbor.right + 1
					//} else {
					//	tree.right = 1
					//}
					////fmt.Fprintf(w, "right: (%s) %d - %d\n", tree.pos(), tree.height, tree.right)
				}
			}
			// ups
			//fmt.Fprintf(w, "ups\n")
			for col := 0; col < cols; col++ {
				for row := 1; row < rows; row++ {
					tree := trees[row][col]
					for i := row - 1; i >= 0; i-- {
						tree.up++
						if trees[i][col].height >= tree.height {
							break
						}
					}
					//neighbor := trees[row-1][col]
					////fmt.Fprintf(w, "\n")
					////fmt.Fprintf(w, "neighbor: (%s) %d - %d\n", neighbor.pos(), neighbor.height, neighbor.up)
					//if tree.height > neighbor.height {
					//	tree.up = neighbor.up + 1
					//} else {
					//	tree.up = 1
					//}
					////fmt.Fprintf(w, "   up: (%s) %d - %d\n", tree.pos(), tree.height, tree.up)
				}
			}
			// downs
			//fmt.Fprintf(w, "downs\n")
			for col := 0; col < cols; col++ {
				for row := rows - 2; row >= 0; row-- {
					tree := trees[row][col]
					for i := row + 1; i < rows; i++ {
						tree.down++
						if trees[i][col].height >= tree.height {
							break
						}
					}
					//neighbor := trees[row+1][col]
					////fmt.Fprintf(w, "\n")
					////fmt.Fprintf(w, "neighbor: (%s) %d - %d\n", neighbor.pos(), neighbor.height, neighbor.down)
					//if tree.height > neighbor.height {
					//	tree.down = neighbor.down + 1
					//} else {
					//	tree.down = 1
					//}
					////fmt.Fprintf(w, " down: (%s) %d - %d\n", tree.pos(), tree.height, tree.down)
				}
			}
			//fmt.Fprintf(w, "details:\n")
			//trees.Format(w, func(t *Tree) string {
			//	return fmt.Sprintf("(%s) %d - up=%d,left=%d,right=%d,down=%d => %d\n",
			//		t.pos(), t.height, t.up, t.left, t.right, t.down, t.score())
			//})

			//fmt.Fprintf(w, "heights:\n")
			//trees.Format(w, func(t *Tree) string { return strconv.Itoa(t.height) })
			//width := 2
			for row := 0; row < rows; row++ {
				//// auto-collapse
				//var expand string
				//var (
				//	start1 int
				//	end1   = cols - 1
				//	start2 int
				//	end2   int
				//)
				//if cols > 20 {
				//	expand = "... "
				//}

				//fmt.Fprintf(w, "   lefts[%d]: %s %s%s\n", row,
				//	formatRow(trees[row], start1, end1, formatLeft), expand, formatRow(trees[row], start2, end2, formatLeft))
				//fmt.Fprintf(w, " heights[%d]: %s %s%s\n", row,
				//	formatRow(trees[row], start1, end1, formatHeight), expand, formatRow(trees[row], start2, end2, formatHeight))
				//fmt.Fprintf(w, "  rights[%d]: %s %s%s\n", row,
				//	formatRow(trees[row], start1, end1, formatRight), expand, formatRow(trees[row], start2, end2, formatRight))
				//fmt.Fprintf(w, "\n")
			}

			//fmt.Fprintf(w, "heights:\n")
			//trees.Format(w, 1, formatHeight)
			//fmt.Fprintf(w, "ups:\n")
			//trees.Format(w, 1, formatUp)
			//fmt.Fprintf(w, "lefts:\n")
			//trees.Format(w, 1, formatLeft)
			//fmt.Fprintf(w, "rights:\n")
			//trees.Format(w, 1, formatRight)
			//fmt.Fprintf(w, "downs:\n")
			//trees.Format(w, 1, formatDown)
			//fmt.Fprintf(w, "scores:\n")
			//trees.Format(w, 1, formatScore)

			// Find the best
			var best *Tree
			trees.Walk(func(t *Tree) {
				if best == nil || best.score() < t.score() {
					best = t
				}
			})
			fmt.Fprintf(w, "best: (%s) %d - up=%d,left=%d,right=%d,down=%d => %d\n",
				best.pos(), best.height, best.up, best.left, best.right, best.down, best.score())
		},
	)
}

type Tree struct {
	height   int
	row, col int
	// If it is visible from each direction
	// TODO(ppg): fill in part 1 with this
	//vLeft, vRight, vUp, vLeft bool

	// View distances
	left, right, up, down int
}

func (t *Tree) pos() string { return fmt.Sprintf("%d,%d", t.col, t.row) }
func (t *Tree) score() int  { return t.up * t.left * t.right * t.down }

var (
	formatUp     = func(t *Tree) string { return fmt.Sprintf("%*d", 2, t.up) }
	formatLeft   = func(t *Tree) string { return fmt.Sprintf("%*d", 2, t.left) }
	formatRight  = func(t *Tree) string { return fmt.Sprintf("%*d", 2, t.right) }
	formatDown   = func(t *Tree) string { return fmt.Sprintf("%*d", 2, t.down) }
	formatHeight = func(t *Tree) string { return fmt.Sprintf("%*d", 2, t.height) }
	formatScore  = func(t *Tree) string { return fmt.Sprintf("%*d", 2, t.score()) }
)

type Matrix [][]*Tree

func (m Matrix) Walk(fn func(*Tree)) {
	for _, row := range m {
		for _, tree := range row {
			fn(tree)
		}
	}
}

func (m Matrix) Format(w io.Writer, gutter int, format func(*Tree) string) {
	for _, row := range m {
		srow := make([]string, len(row))
		for j, tree := range row {
			srow[j] = format(tree)
		}
		fmt.Fprintf(w, "%s\n", strings.Join(srow, strings.Repeat(" ", gutter)))
	}
}

func formatRow(in []*Tree, start, end int, format func(*Tree) string) string {
	out := make([]string, end-start)
	for i, t := range in[start:end] {
		out[i] = fmt.Sprintf("%s", format(t))
	}
	return strings.Join(out, " ")
}
