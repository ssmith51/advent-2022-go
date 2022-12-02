package main

import (
	"advent-01/internal/pgk/utils"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var INPUT_FILE = "test.txt"

type Elf struct {
	Food      []int
	TotalFood int
}

func load(filename string) map[int]Elf {

	fi, err := os.Open(filename)
	utils.HandleError(err)
	defer fi.Close()

	//Hold all the elves and their food
	elves := make(map[int]Elf)

	//Set the first elf as #1 in case we need it
	var elfId = 1

	//Read in the file line by line
	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "")

		if len(line) == 0 {
			elfId += 1
		} else {
			val, err := strconv.Atoi(line)
			utils.HandleError(err)

			elf := elves[elfId]
			elf.Food = append(elf.Food, val)
			elves[elfId] = elf
		}
	}

	return elves
}

func sumCalories(elves map[int]Elf) map[int]Elf {

	for i, elf := range elves {
		log.Println(elf)

		calores := 0
		for _, cal := range elf.Food {
			calores += cal
		}
		elf.TotalFood = calores
		elves[i] = elf
	}

	return elves
}

func puzzle_1(elves map[int]Elf) {
	log.Println(elves)
}

func main() {
	log.Println("Starting Day 01")
	elves := load(INPUT_FILE)
	elves = sumCalories(elves)
	puzzle_1(elves)
}
