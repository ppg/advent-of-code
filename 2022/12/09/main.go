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

var parser = func(line string) *Move {
	return &Move{
		dir:   Direction(line[0:1]),
		steps: framework.ParseInt(line[2:]),
	}
}

func solution0(w io.Writer, runner *framework.Runner[*Move]) {
	xMin := parseEnvInt("X_MIN")
	xMax := parseEnvInt("X_MAX")
	yMin := parseEnvInt("Y_MIN")
	yMax := parseEnvInt("Y_MAX")
	var showDebug bool
	if xMax == 0 {
		showDebug = true
	}

	var knotCount int
	runner.ByPart(
		func() { knotCount = 2 },
		func() { knotCount = 10 },
	)
	knots := make([]*Position, knotCount)
	for i := range knots {
		label := fmt.Sprintf("%d", i)
		if i == 0 {
			label = "H"
		}
		knots[i] = &Position{label: label}
	}
	draw(w, showDebug, xMin, xMax, yMin, yMax, knots...)
	visited := make([]*Position, 0, 1000)
	mVisited := make(map[int]map[int]bool)
	//visited := make(map[Position]bool)
	for move := range runner.Lines() {
		if !showDebug {
			fmt.Fprintf(w, "\n== %s %d ==\n", move.dir, move.steps)
			fmt.Fprintln(w)
		}
		for i := 0; i < move.steps; i++ {
			// move head
			knots[0].step(move.dir)

			// move each knot in order
			for i := 1; i < len(knots); i++ {
				moveTail(w, i, knots[i-1], knots[i])
			}

			//// move tail if necessary
			//tail := knots[len(knots)-1]
			//deltaX := head.x - tail.x
			//deltaY := head.y - tail.y
			//tailStep := tailSteps[deltaX][deltaY]
			////fmt.Fprintf(w, "head=%s\n", head)
			////fmt.Fprintf(w, "tail=%s\n", tail)
			//fmt.Fprintf(w, "head=%s tail=%s deltaX=(%d,%d) => %s\n", head, tail, deltaX, deltaY, tailStep)
			//tail.step(tailStep)

			// Draw after steps
			//draw(w, showDebug, xMin, xMax, yMin, yMax, knots...)

			// Check what the tail has visited
			tail := knots[len(knots)-1]
			if !mVisited[tail.x][tail.y] {
				if _, ok := mVisited[tail.x]; !ok {
					mVisited[tail.x] = make(map[int]bool)
				}
				mVisited[tail.x][tail.y] = true
				visited = append(visited, &Position{label: "#", x: tail.x, y: tail.y})
			}
			for i := 0; i < len(knots); i++ {
				xMin = min(xMin, knots[i].x)
				xMax = max(xMax, knots[i].x)
				yMin = min(yMin, knots[i].y)
				yMax = max(yMax, knots[i].y)
			}
		}

		// Draw after moves
		draw(w, showDebug, xMin, xMax, yMin, yMax, knots...)
	}

	fmt.Fprintln(w, "== Visited ==")
	draw(w, showDebug, xMin, xMax, yMin, yMax, visited...)
	fmt.Fprintf(w, "Count: %d\n", len(visited))
	if showDebug {
		fmt.Fprintln(w, "For debug:")
		runner.ByPart(
			func() {
				fmt.Fprintf(w, "X_MIN=%d X_MAX=%d Y_MIN=%d Y_MAX=%d PART=%d go run main.go %s\n", xMin, xMax, yMin, yMax, 1, strings.Join(os.Args[1:], " "))
			},
			func() {
				fmt.Fprintf(w, "X_MIN=%d X_MAX=%d Y_MIN=%d Y_MAX=%d PART=%d go run main.go %s\n", xMin, xMax, yMin, yMax, 2, strings.Join(os.Args[1:], " "))
			},
		)
	}
}

func moveTail(w io.Writer, i int, prev, knot *Position) {
	deltaX := prev.x - knot.x
	deltaY := prev.y - knot.y

	var knotStep Direction
	switch {
	case deltaX == 0 && deltaY == 0: // no movement

	case deltaX == 0 && deltaY > 1:
		knotStep = Up
	case deltaX == 0 && deltaY < -1:
		knotStep = Down
	case deltaX > 1 && deltaY == 0:
		knotStep = Right
	case deltaX < -1 && deltaY == 0:
		knotStep = Left

	case (deltaX > 1 && deltaY > 0) || (deltaX > 0 && deltaY > 1):
		knotStep = UpRight
	case (deltaX < -1 && deltaY > 0) || (deltaX < 0 && deltaY > 1):
		knotStep = UpLeft
	case (deltaX < -1 && deltaY < 0) || (deltaX < 0 && deltaY < -1):
		knotStep = DownLeft
	case (deltaX > 1 && deltaY < 0) || (deltaX > 0 && deltaY < -1):
		knotStep = DownRight
	}

	//fmt.Fprintf(w, "%d: prev=%s knot=%s deltaX=(%d,%d) => %s\n", i, prev, knot, deltaX, deltaY, knotStep)
	knot.step(knotStep)
}

func moveTailBackup(w io.Writer, i int, prev, knot *Position) {
	deltaX := prev.x - knot.x
	deltaY := prev.y - knot.y
	knotStep := knotSteps[deltaX][deltaY]
	fmt.Fprintf(w, "%d: prev=%s knot=%s deltaX=(%d,%d) => %s\n", i, prev, knot, deltaX, deltaY, knotStep)
	knot.step(knotStep)
}

// knotSteps [delta x][delta y]
// delta is from the knot ahead, so - means left/down
var knotSteps = map[int]map[int]Direction{
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

func draw(w io.Writer, showDebug bool, xMin, xMax, yMin, yMax int, positions ...*Position) {
	// If we're going to show debug, skip printing
	if showDebug {
		return
	}
	for y := yMax; y >= yMin; y-- {
		for x := xMin; x < xMax; x++ {
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
	fmt.Fprintln(w)
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

// TODO(ppg): move to framework

func parseEnvInt(key string) int {
	if value, ok := os.LookupEnv(key); ok {
		return framework.ParseInt(value)
	}
	return 0
}

func min(values ...int) int {
	var out int
	for _, value := range values {
		if value < out {
			out = value
		}
	}
	return out
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
