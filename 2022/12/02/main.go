package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	framework "github.com/ppg/advent-of-code/2022/12/framework"
)

func main() {
	framework.Register(solution0)
	framework.Register(solution1)
	framework.Run(os.Stdout)
}

// Solution 0 matches the columns to the play, and then calculates the score.
func solution0(w io.Writer, runner *framework.Runner) {
	var total int
	fmt.Fprintf(w, "%10s %10s %5s\n", "opponent", "self", "score")
	for runner.Scan() {
		// Get the entries from the line
		line := strings.TrimSpace(runner.Text())
		cols := strings.Split(line, " ")
		if len(cols) != 2 {
			panic(fmt.Errorf("unexpected input: %s", line))
		}

		// The opponent's play is always the first column
		opponent := opponentPlays[cols[0]]

		// Get the self play depending on which question part
		var self Play
		runner.ByPart(
			// In part 1 the second column is interpreted as the self play, so lookup
			// the self play in it's map
			func() {
				self = selfPlays[cols[1]]
			},
			// in part 2 the second column is interpreted as the desired outcome, so
			// lookup the self play in a map of outcomes from desired outcome to what
			// the opponent has
			func() {
				self = outcomePlays[cols[1]][opponent]
			},
		)

		// Score self play and the outcome and accumulate
		score := playScores[self] + vsScores[self][opponent]
		total += score

		fmt.Fprintf(w, "%10s %10s %5d\n", opponent, self, score)
	}
	fmt.Fprintf(w, "Total: %d\n", total)
}

type Play string

const (
	Rock     = "rock"
	Paper    = "paper"
	Scissors = "scissors"
)

var opponentPlays = map[string]Play{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
}
var selfPlays = map[string]Play{
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

// The Elf finishes helping with the tent and sneaks back over to you. "Anyway,
// the second column says how the round needs to end: X means you need to lose,
// Y means you need to end the round in a draw, and Z means you need to win.
// Good luck!"
var outcomePlays = map[string]map[Play]Play{
	"X": map[Play]Play{ // lose
		Rock:     Scissors,
		Paper:    Rock,
		Scissors: Paper,
	},
	"Y": map[Play]Play{ // draw
		Rock:     Rock,
		Paper:    Paper,
		Scissors: Scissors,
	},
	"Z": map[Play]Play{ // win
		Rock:     Paper,
		Paper:    Scissors,
		Scissors: Rock,
	},
}

var playScores = map[Play]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

// lookup with self play, opponent play
var vsScores = map[Play]map[Play]int{
	Rock: map[Play]int{
		Rock:     3, // draw
		Paper:    0, // loss
		Scissors: 6, // win
	},
	Paper: map[Play]int{
		Rock:     6, // win
		Paper:    3, // draw
		Scissors: 0, // loss
	},
	Scissors: map[Play]int{
		Rock:     0, // loss
		Paper:    6, // win
		Scissors: 3, // draw
	},
}

// Solution 1 looks up the pre-computed score given the column inputs
func solution1(w io.Writer, runner *framework.Runner) {
	var total int
	for runner.Scan() {
		// Get the entries from the line
		line := strings.TrimSpace(runner.Text())
		cols := strings.Split(line, " ")
		if len(cols) != 2 {
			panic(fmt.Errorf("unexpected input: %s", line))
		}

		var score int
		runner.ByPart(
			func() { score = part1Scores[cols[0]][cols[1]] },
			func() { score = part2Scores[cols[0]][cols[1]] },
		)
		total += score
	}
	fmt.Fprintf(w, "Total: %d\n", total)
}

var part1Scores = map[string]map[string]int{
	"A": map[string]int{ // opponent: rock
		"X": 4, // self: rock(1) + draw(3)
		"Y": 8, // self: paper(2) + win(6)
		"Z": 3, // self: scissors(3) + loss(0)
	},
	"B": map[string]int{ // opponent: paper
		"X": 1, // self: rock(1) + loss(0)
		"Y": 5, // self: paper(2) + draw(3)
		"Z": 9, // self: scissors(3) + win(6)
	},
	"C": map[string]int{ // opponent: scissors
		"X": 7, // self: rock(1) + win(6)
		"Y": 2, // self: paper(2) + loss(0)
		"Z": 6, // self: scissors(3) + draw(3)
	},
}
var part2Scores = map[string]map[string]int{
	"A": map[string]int{ // opponent: rock
		"X": 3, // loss(0) + scissors(3)
		"Y": 4, // draw(3) + rock(1)
		"Z": 8, // win(6) + paper(2)
	},
	"B": map[string]int{ // opponent: paper
		"X": 1, // loss(0) + rock(1)
		"Y": 5, // draw(3) + paper(2)
		"Z": 9, // win(6) + scissors(3)
	},
	"C": map[string]int{ // opponent: scissors
		"X": 2, // loss(0) + paper(2)
		"Y": 6, // draw(3) + scissors(3)
		"Z": 7, // win(6) + rock(1)
	},
}
