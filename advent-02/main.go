package main

import (
	"advent-02/internal/pgk/utils"
	"bufio"
	"log"
	"os"
	"strings"
)

const (
	Rock          string = "A"
	Paper         string = "B"
	Scissors      string = "C"
	ScoreRock     int    = 1
	ScorePaper    int    = 2
	ScoreScissors int    = 3
	PlayRock      string = "X"
	PlayScissors  string = "Z"
	PlayPaper     string = "Y"
	Loss          int    = 0
	Draw          int    = 3
	Win           int    = 6
	PlayWin       string = "Z"
	PlayDraw      string = "Y"
	PlayLose      string = "X"
)

type Round struct {
	Opponent string
	Player   string
}

func load(filename string) []Round {
	fi, err := os.Open(filename)
	utils.HandleError(err)
	defer fi.Close()

	game := []Round{}

	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "")
		plays := strings.Split(line, " ")
		r := Round{}
		r.Opponent = plays[0]
		r.Player = plays[1]
		game = append(game, r)

	}

	return game
}

func puzzle_1(game []Round) {

	totalScore := 0
	for _, round := range game {

		//Opponent Plays Rock
		if round.Opponent == Rock {
			switch round.Player {
			case PlayPaper:
				totalScore += Win
				totalScore += ScorePaper

			case PlayRock:
				totalScore += Draw
				totalScore += ScoreRock

			case PlayScissors:
				totalScore += ScoreScissors
			}

			//Opponent Plays Paper
		} else if round.Opponent == Paper {
			switch round.Player {
			case PlayScissors:
				totalScore += Win
				totalScore += ScoreScissors

			case PlayPaper:
				totalScore += Draw
				totalScore += ScorePaper

			case PlayRock:
				totalScore += ScoreRock

			}

			//Opponent Plays Scissors
		} else if round.Opponent == Scissors {
			switch round.Player {
			case PlayRock:
				totalScore += Win
				totalScore += ScoreRock

			case PlayScissors:
				totalScore += Draw
				totalScore += ScoreScissors

			case PlayPaper:
				totalScore += ScorePaper

			}

		}

	}

	log.Printf("Puzzle 1 Total Score: %d", totalScore)
}

func puzzle_2(game []Round) {

	totalScore := 0
	for _, round := range game {

		if round.Player == PlayWin {
			totalScore += Win
			switch round.Opponent {
			case Rock:
				totalScore += ScorePaper

			case Paper:
				totalScore += ScoreScissors

			case Scissors:
				totalScore += ScoreRock

			}

		} else if round.Player == PlayDraw {
			totalScore += Draw
			switch round.Opponent {
			case Rock:
				totalScore += ScoreRock

			case Paper:
				totalScore += ScorePaper

			case Scissors:
				totalScore += ScoreScissors

			}

		} else if round.Player == PlayLose {
			totalScore += Loss
			switch round.Opponent {
			case Rock:
				totalScore += ScoreScissors

			case Paper:
				totalScore += ScoreRock

			case Scissors:
				totalScore += ScorePaper

			}

		}

	}

	log.Printf("Puzzle 2 Total Score: %d", totalScore)
}

func main() {

	game := load("input.txt")
	puzzle_1(game)
	puzzle_2(game)
}
