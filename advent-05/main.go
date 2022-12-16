package main

import (
	"advent-05/internal/pgk/utils"
	"bufio"
	"log"
	"os"
)

func load(filename string) map[int][]float64 {

	groups := map[int][]float64{}

	fi, err := os.Open(filename)
	utils.HandleError(err)
	defer fi.Close()

	// lines := []string{}

	scanner := bufio.NewScanner(fi)

	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		log.Println(string(runes))
	}

	return groups
}

func main() {
	log.Println("Advent of Code Day 05")
	load("test.txt")
}
