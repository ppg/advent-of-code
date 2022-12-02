package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := "input.txt"
	if len(os.Args) > 1 {
		input = os.Args[1]
	}
	readFile, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	var total int

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	fmt.Printf("%10s %10s %5s\n", "opponent", "self", "score")
	for fileScanner.Scan() {
		// Get the entries from the line
		line := strings.TrimSpace(fileScanner.Text())
		cols := strings.Split(line, " ")
		if len(cols) != 2 {
			panic(fmt.Errorf("unexpected input: %s", line))
		}
		// The opponent's play is always the first column
		opponent := opponentPlays[cols[0]]

		// Get the self play depending on which question part
		var self Play
		if os.Getenv("PART") == "2" {
			// in part 2 the second column is interpreted as the desired outcome, so
			// lookup the self play in a map of outcomes from desired outcome to what
			// the opponent has
			self = outcomePlays[cols[1]][opponent]
		} else {
			// in part 1 the second column is interpreted as the self play, so lookup the self play in it's map
			self = selfPlays[cols[1]]
		}

		// Score self play and the outcome and accumulate
		score := playScores[self] + vsScores[self][opponent]
		total += score

		fmt.Printf("%10s %10s %5d\n", opponent, self, score)
	}
	fmt.Printf("Total: %d\n", total)
}

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

type Play string

const (
	Rock     = "rock"
	Paper    = "paper"
	Scissors = "scissors"
)
