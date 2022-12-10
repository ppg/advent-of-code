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

var parser = func(line string) *Move {
	return &Move{
		dir:   Direction(line[0:1]),
		steps: framework.ParseInt(line[2:]),
	}
}

func solution0(w io.Writer, runner *framework.Runner[*Move]) {
	// R 4
	// U 4
	// L 3
	// D 1
	// R 4
	// D 1
	// L 5
	// R 2
	head := &Position{label: "H"}
	tail := &Position{label: "T"}
	draw(w, head, tail)
	fmt.Fprintln(w)
	visited := make([]*Position, 0, 1000)
	mVisited := make(map[int]map[int]bool)
	//visited := make(map[Position]bool)
	for move := range runner.Lines() {
		fmt.Fprintf(w, "\n== %s %d ==\n", move.dir, move.steps)
		fmt.Fprintln(w)
		for i := 0; i < move.steps; i++ {
			// move head
			head.step(move.dir)

			// move tail if necessary
			deltaX := head.x - tail.x
			deltaY := head.y - tail.y
			tailStep := tailSteps[deltaX][deltaY]
			//fmt.Fprintf(w, "head=%s\n", head)
			//fmt.Fprintf(w, "tail=%s\n", tail)
			fmt.Fprintf(w, "head=%s tail=%s deltaX=(%d,%d) => %s\n", head, tail, deltaX, deltaY, tailStep)
			tail.step(tailStep)

			// Draw after moves
			draw(w, head, tail)
			fmt.Fprintln(w)

			if !mVisited[tail.x][tail.y] {
				if _, ok := mVisited[tail.x]; !ok {
					mVisited[tail.x] = make(map[int]bool)
				}
				mVisited[tail.x][tail.y] = true
				visited = append(visited, &Position{label: "#", x: tail.x, y: tail.y})
			}
		}
	}

	fmt.Fprintln(w, "== Visited ==")
	draw(w, visited...)
	fmt.Fprintf(w, "Count: %d\n", len(visited))
}

// tailSteps [delta x][delta y]
// delta -2 to 2 for ta, no -2 in both
var tailSteps = map[int]map[int]Direction{
	-2: map[int]Direction{
		-1: DownLeft, // head is 2 left 1 down
		0:  Left,     // head is 2 left
		1:  UpLeft,   // head is 2 left 1 up
	},
	-1: map[int]Direction{
		-2: DownLeft, // head is 1 left 2 down
		//-1: Down or Left, // head is 1 left 1 down
		//0: Left, // head is 1 left
		//1:  Up or Left,   // head is 1 left 1 up
		2: UpLeft, // head is 1 left 2 up
	},
	0: map[int]Direction{
		-2: Down, // head is 1 down
		//-1: Down, // head is 1 down
		// 0:  Left,     // head is at tail
		//1:  Left,   // head is 1 up
		2: Up, // head is 1 up
	},
	1: map[int]Direction{
		-2: DownRight, // head is 1 right 2 down
		//-1: Down or Right, // head is 1 right 1 down
		//0: Right, // head is 1 right
		//1:  Up or Right,   // head is 1 right 1 up
		2: UpRight, // head is 1 right 2 up
	},
	2: map[int]Direction{
		-1: DownRight, // head is 2 right 1 down
		0:  Right,     // head is 2 right
		1:  UpRight,   // head is 2 right 1 up
	},
}

var start = &Position{label: "s"}

func draw(w io.Writer, positions ...*Position) {
	var xSize int = 6
	var ySize int = 5
	for _, pos := range positions {
		if pos.x > xSize {
			xSize = pos.x
		}
		if pos.y > ySize {
			ySize = pos.y
		}
	}
	for y := ySize - 1; y >= 0; y-- {
		for x := 0; x < xSize; x++ {
			var found bool
			for _, pos := range positions {
				if pos.at(x, y) {
					fmt.Fprintf(w, pos.label)
					found = true
					break
				}
			}
			if !found {
				// nothing at the position, .
				if start.at(x, y) {
					fmt.Fprintf(w, "s")
				} else {
					fmt.Fprintf(w, ".")
				}
			}
		}
		fmt.Fprintln(w)
	}
}

func max(values ...int) int {
	var out int
	for _, value := range values {
		if value > out {
			out = value
		}
	}
	return out
}

type Position struct {
	label string
	x, y  int
}

func (p Position) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func (p Position) at(x, y int) bool {
	return p.x == x && p.y == y
}

func (p *Position) step(dir Direction) {
	switch dir {
	case Right:
		p.x++
	case Left:
		p.x--
	case Up:
		p.y++
	case Down:
		p.y--
	case UpRight:
		p.step(Up)
		p.step(Right)
	case UpLeft:
		p.step(Up)
		p.step(Left)
	case DownRight:
		p.step(Down)
		p.step(Right)
	case DownLeft:
		p.step(Down)
		p.step(Left)
	}
}

type Move struct {
	dir   Direction
	steps int
}

type Direction string

const (
	Right Direction = "R"
	Left  Direction = "L"
	Up    Direction = "U"
	Down  Direction = "D"

	UpRight   Direction = "UR"
	UpLeft    Direction = "UL"
	DownRight Direction = "DR"
	DownLeft  Direction = "DL"
)
