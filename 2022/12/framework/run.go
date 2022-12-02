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

type SolutionHandler[T any] func(io.Writer, *Runner[T])
type SolutionParser[T any] func(string) T

func Register[T any](parse SolutionParser[T], handle SolutionHandler[T]) {
	solutions = append(solutions, &solution[T]{handle: handle, parse: parse})
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
	s := solutions[index]

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
	s.Handle(w, scanner, part)
}

type Runner[T any] struct {
	scanner *bufio.Scanner
	part    int
	parse   SolutionParser[T]
}

func (r *Runner[T]) Lines() <-chan T {
	ch := make(chan T)
	go func() {
		for r.scanner.Scan() {
			ch <- r.parse(r.scanner.Text())
		}
		close(ch)
	}()
	return ch
}

func (r *Runner[T]) ByPart(part1, part2 func()) {
	switch r.part {
	case 1:
		part1()
	case 2:
		part2()
	default:
		panic(fmt.Errorf("invalid part: %d", r.part))
	}
}

var solutions []isolution

type isolution interface {
	Handle(io.Writer, *bufio.Scanner, int)
}

type solution[T any] struct {
	handle SolutionHandler[T]
	parse  SolutionParser[T]
}

func (s *solution[T]) Handle(w io.Writer, scanner *bufio.Scanner, part int) {
	s.handle(w, &Runner[T]{scanner: scanner, part: part, parse: s.parse})
}
