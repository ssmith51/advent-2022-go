package main

import (
	"advent-05/internal/pgk/utils"
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Move struct {
	Amount int
	From   int
	To     int
}

func load(filename string) (map[int][]string, []Move) {

	fi, err := os.Open(filename)
	utils.HandleError(err)
	defer fi.Close()

	scanner := bufio.NewScanner(fi)

	stacks := make(map[int][]string)
	moves := []Move{}

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			if line[0] == '[' || line[0] == ' ' && line[1] != '1' {
				stacks = parseCrate(line, stacks)
			} else if line[0] == 'm' {
				moves = append(moves, parseMove(line))
			}
		}

	}

	return stacks, moves
}

func parseCrate(line string, stacks map[int][]string) map[int][]string {

	i := 3
	x := 0
	stackNum := 1
	for i <= len(line) {
		crate := line[x:i]
		stack := stackNum

		if string(crate[1]) != " " {
			stacks[stack] = append(stacks[stack], string(crate[1]))
		}

		x += 4
		i += 4
		stackNum++
	}

	return stacks

}

func parseMove(line string) Move {
	r, _ := regexp.Compile(`move\s(\d+)\sfrom\s(\d+)\sto\s(\d+)`)
	matches := r.FindStringSubmatch(line)
	move := Move{}
	m, _ := strconv.Atoi(matches[1])
	f, _ := strconv.Atoi(matches[2])
	t, _ := strconv.Atoi(matches[3])
	move.Amount = m
	move.From = f
	move.To = t

	return move
}

func puzzle_1(stacks map[int][]string, moves []Move) {

	log.Println("--- Puzzle 1 ---")

	for _, move := range moves {
		i := 1
		for i <= move.Amount {

			var newToStack []string

			newToStack = append(newToStack, stacks[move.From][0])
			newToStack = append(newToStack, stacks[move.To][:]...)
			stacks[move.To] = newToStack
			stacks[move.From] = stacks[move.From][1:len(stacks[move.From])]
			i++
		}
	}

	var topRow []string
	i := 1
	for i <= len(stacks) {
		topRow = append(topRow, stacks[i][0])
		i++
	}

	log.Printf("Top Row: %s", topRow)
}

func puzzle_2(stacks map[int][]string, moves []Move) {

	log.Println("--- Puzzle 2 ---")
	for _, move := range moves {
		// i := 1
		// for i <= move.Amount {

		var newToStack []string

		newToStack = append(newToStack, stacks[move.From][0:move.Amount]...)
		newToStack = append(newToStack, stacks[move.To][:]...)
		stacks[move.To] = newToStack
		stacks[move.From] = stacks[move.From][move.Amount:len(stacks[move.From])]
		// i++
		// }
	}

	var topRow []string
	i := 1
	for i <= len(stacks) {
		topRow = append(topRow, stacks[i][0])
		i++
	}

	log.Printf("Top Row: %s", topRow)

}

func main() {
	log.Println("Advent of Code Day 05")
	stacks, moves := load("input.txt")
	// puzzle_1(stacks, moves)
	puzzle_2(stacks, moves)

}
