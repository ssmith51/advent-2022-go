package main

import (
	"advent-03/internal/pgk/utils"
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"
)

func load(filename string) []string {
	fi, err := os.Open(filename)
	utils.HandleError(err)
	defer fi.Close()

	lines := []string{}

	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "")
		lines = append(lines, line)
	}

	return lines
}

func contains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}

func puzzle_1(rucksacks []string) {

	validation := map[string]int{}
	lowerOffset := 96
	upperOffset := 38
	sum := 0
	for _, rucksack := range rucksacks {

		length := len(rucksack)

		first := rucksack[0:(length / 2)]
		second := rucksack[length/2:]

		sortedFirst := strings.Split(first, "")
		sort.Strings(sortedFirst)

		sortedSecond := strings.Split(second, "")
		sort.Strings(sortedSecond)

		itemTypes := []string{}

		for _, a := range sortedFirst {

			if contains(sortedSecond, a) {

				sort.Strings(itemTypes)
				if !contains(itemTypes, a) {
					itemTypes = append(itemTypes, a)

					c := []rune(a)[0]

					if unicode.IsUpper(c) {
						sum += (int(c) - upperOffset)
						validation[a] = (int(c) - upperOffset)
					} else {
						sum += (int(c) - lowerOffset)
						validation[a] = (int(c) - lowerOffset)
					}

				}

			}
		}
	}
	log.Printf("Total Sum of Item Proirities: %d", sum)
}

func puzzle_2(rucksacks []string) {

	lowerOffset := 96
	upperOffset := 38
	sum := 0

	for i := range rucksacks {

		if (i % 3) == 0 {

			first := rucksacks[i]
			second := rucksacks[i+1]
			third := rucksacks[i+2]

			//Sort each list
			sortedFirst := strings.Split(first, "")
			sort.Strings(sortedFirst)

			sortedSecond := strings.Split(second, "")
			sort.Strings(sortedSecond)

			sortedThird := strings.Split(third, "")
			sort.Strings(sortedThird)

			itemTypes := []string{}
			for _, a := range sortedFirst {

				//Check if the badge is found in all 3 rucksacks
				if contains(sortedSecond, a) && contains(sortedThird, a) {

					//Identify if the badge is already accounted for
					sort.Strings(itemTypes)
					if !contains(itemTypes, a) {
						itemTypes = append(itemTypes, a)

						c := []rune(a)[0]

						//Determine if upper or lower case
						if unicode.IsUpper(c) {
							sum += (int(c) - upperOffset)
						} else {
							sum += (int(c) - lowerOffset)
						}

					}

				}
			}

		}
	}

	log.Printf("Total Sum of Badges: %d", sum)

}

func main() {
	log.Println("Advent 2022 Day 03")
	rucksacks := load("input.txt")
	puzzle_1(rucksacks)
	puzzle_2(rucksacks)
}
