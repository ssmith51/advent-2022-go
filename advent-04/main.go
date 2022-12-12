package main

import (
	"advent-04/internal/pgk/utils"
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/montanaflynn/stats"
)

func load(filename string) map[int][]float64 {

	groups := map[int][]float64{}

	fi, err := os.Open(filename)
	utils.HandleError(err)
	defer fi.Close()

	// lines := []string{}

	scanner := bufio.NewScanner(fi)

	i := 0
	for scanner.Scan() {

		line := strings.Trim(scanner.Text(), "")

		r, _ := regexp.Compile(`(\d.*)-(\d.*),(\d.*)-(\d.*)`)
		parsed := r.FindStringSubmatch(line)

		a, _ := strconv.ParseFloat(parsed[1], 64)
		b, _ := strconv.ParseFloat(parsed[2], 64)
		c, _ := strconv.ParseFloat(parsed[3], 64)
		d, _ := strconv.ParseFloat(parsed[4], 64)

		values := []float64{a, b, c, d}

		groups[i] = values

		i++
	}

	return groups
}

func puzzle_1(groups map[int][]float64) {
	overlaps := 0
	for _, val := range groups {

		isOverlap := calculateFullOverlap(val[0], val[1], val[2], val[3])

		if !isOverlap {
			isOverlap = calculateFullOverlap(val[2], val[3], val[0], val[1])
		}

		if isOverlap {
			overlaps++
		}

	}

	log.Println("Overlaps: ", overlaps)

}

func puzzle_2(groups map[int][]float64) {
	overlaps := 0
	for _, val := range groups {

		isOverlap := calculatePartialOverlap(val[0], val[1], val[2], val[3])

		if !isOverlap {
			isOverlap = calculatePartialOverlap(val[2], val[3], val[0], val[1])
		}

		if isOverlap {
			overlaps++
		}

	}

	log.Println("Overlaps: ", overlaps)

}

func calculatePartialOverlap(a float64, b float64, c float64, d float64) bool {

	isOverlap := false

	fElf := []float64{a, b + 1, d + 1}
	sElf := []float64{a, b + 1, c}
	fMedian, _ := stats.Median(fElf)
	sMedian, _ := stats.Median(sElf)

	firstElfOverlap := fMedian - sMedian

	if firstElfOverlap > 0 {
		isOverlap = true
	}
	return isOverlap
}

func calculateFullOverlap(a float64, b float64, c float64, d float64) bool {

	isOverlap := false

	fElf := []float64{a, b + 1, d + 1}
	sElf := []float64{a, b + 1, c}
	fMedian, _ := stats.Median(fElf)
	sMedian, _ := stats.Median(sElf)

	firstElfOverlap := fMedian - sMedian

	if (d-c)+1 <= firstElfOverlap {
		isOverlap = true
	}
	return isOverlap
}

func main() {
	log.Println("Advent Day 04")
	groups := load("input.txt")
	puzzle_1(groups)
	puzzle_2(groups)
}
