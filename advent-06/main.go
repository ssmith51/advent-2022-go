package main

import (
	"advent-06/internal/pgk/utils"
	"bufio"
	"log"
	"os"
)

func load(filename string) []rune {

	fi, err := os.Open(filename)
	utils.HandleError(err)
	defer fi.Close()

	scanner := bufio.NewScanner(fi)
	runes := []rune{}
	for scanner.Scan() {
		line := scanner.Text()
		runes = []rune(line)
	}

	return runes

}

func puzzle_1(markers []rune) {

	packets := []rune{}
	processed := 0
	start := false
	msgLen := 4

	for processed < len(markers) && !start {

		packets = append(packets, markers[processed])

		if len(packets) > msgLen {
			packets = packets[1:]
			startPacket := make(map[rune]string)
			for _, p := range packets {
				startPacket[p] = string(p)
			}

			if len(startPacket) == msgLen {
				start = true
			}
		}
		processed++
	}

	log.Printf("Start Packet Discovered after %d signals", processed)

}

func puzzle_2(markers []rune) {

	packets := []rune{}
	processed := 0
	start := false
	msgLen := 14

	for processed < len(markers) && !start {

		packets = append(packets, markers[processed])

		if len(packets) > msgLen {
			packets = packets[1:]
			startPacket := make(map[rune]string)
			for _, p := range packets {
				startPacket[p] = string(p)
			}

			if len(startPacket) == msgLen {
				start = true
			}
		}
		processed++
	}

	log.Printf("Start Packet Discovered after %d signals", processed)

}

func main() {
	log.Println("Advent of Code Day 06")
	markers := load("input.txt")
	puzzle_1(markers)
	puzzle_2(markers)
}
