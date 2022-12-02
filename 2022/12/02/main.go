package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	var total int

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	fmt.Printf("%10s %10s %5s\n", "opponent", "self", "score")
	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())
		plays := strings.Split(line, " ")
		if len(plays) != 2 {
			panic(fmt.Errorf("unexpected input: %s", line))
		}
		opponent := opponentPlays[plays[0]]
		self := selfPlays[plays[1]]
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
