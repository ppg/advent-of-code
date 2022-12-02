package framework

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	EnvPart     = "PART"
	EnvSolution = "SOLUTION"
)

var solutions []Solution

func Register(solution Solution) {
	solutions = append(solutions, solution)
}

func Run(w io.Writer) {
	// Parse solution input
	index := len(solutions) - 1
	if sIndex := os.Getenv(EnvSolution); sIndex != "" {
		var err error
		index, err = strconv.Atoi(sIndex)
		if err != nil {
			panic(fmt.Errorf("solution index invalid: %s", sIndex))
		}
		if index >= len(solutions) {
			panic(fmt.Errorf("solution index out of range: %d (max:%d)", index, len(solutions)-1))
		}
	}
	solution := solutions[index]

	// Parse the part input
	var part int
	sPart := os.Getenv(EnvPart)
	switch sPart {
	case "", "2":
		part = 2
	case "1":
		part = 1
	default:
		panic(fmt.Errorf("unrecognized part: %s", sPart))
	}

	// Read the input
	// Read the input
	file, close := readInput(w)
	defer close()

	// Create a scanner
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Run the solution
	fmt.Fprintf(w, "running with solution=%d, part=%d\n", index, part)
	solution(w, &Runner{scanner: scanner, part: part})
}

type Runner struct {
	scanner *bufio.Scanner
	part    int
}

func (r *Runner) Scan() bool   { return r.scanner.Scan() }
func (r *Runner) Text() string { return r.scanner.Text() }

func (r *Runner) ByPart(part1, part2 func()) {
	switch r.part {
	case 1:
		part1()
	case 2:
		part2()
	default:
		panic(fmt.Errorf("invalid part: %d", r.part))
	}
}

type Solution func(io.Writer, *Runner)

type LineScanner interface {
	Scan() bool
	Text() string
}
