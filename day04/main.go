package main

import (
	"aoc/lib"
	"bufio"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	partOneAnswer := partOne()
	partTwoAnswer := partTwo()

	lib.LogAnswers(partOneAnswer, partTwoAnswer)
}

func getCards() ([][]int, [][]int) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	lines := bufio.NewScanner(file)

	var allWinningNumbers [][]int
	var allElfNumbers [][]int

	for lines.Scan() {
		line := lines.Text()

		winningNumbers, elfNumbers := parseLine(line)

		allWinningNumbers = append(allWinningNumbers, winningNumbers)
		allElfNumbers = append(allElfNumbers, elfNumbers)
	}

	return allWinningNumbers, allElfNumbers
}

func parseLine(line string) ([]int, []int) {
	numbers := strings.Split(line, ":")[1]
	winningNumbers, elfNumbers, success := strings.Cut(numbers, "|")
	if !success {
		log.Fatalf("Could not parse line: %s", numbers)
	}

	return getNumsFromString(winningNumbers), getNumsFromString(elfNumbers)
}

func getNumsFromString(line string) []int {
	var out []int

	line = strings.Trim(line, " ")
	nums := lib.Filter(strings.Split(line, " "), func(item string) bool {
		return strings.Trim(item, " ") != ""
	})

	for _, num := range nums {
		asInt, err := strconv.Atoi(strings.Trim(num, " "))
		if err != nil {
			log.Fatalf("Failed to convert to int: %s", num)
		}

		out = append(out, asInt)
	}

	return out
}

func getCardScore(numWinningCards int) int {
	if numWinningCards == 0 {
		return 0
	}

	return int(math.Pow(2, float64(numWinningCards-1)))
}

func getNumWinningCards(winningCard, elfCard []int) int {
	slices.Sort(winningCard)
	slices.Sort(elfCard)

	numWinners := 0

	for _, v := range elfCard {
		_, found := slices.BinarySearch(winningCard, v)

		if found {
			numWinners++
		}
	}

	return numWinners
}

func partOne() int {
	winningNumbers, elfNumbers := getCards()
	score := 0

	for i := 0; i < len(winningNumbers); i++ {
		winningCard := winningNumbers[i]
		elfCard := elfNumbers[i]

		slices.Sort(winningCard)
		slices.Sort(elfCard)

		numWinners := getNumWinningCards(winningCard, elfCard)

		score += getCardScore(numWinners)
	}

	return score
}

func partTwo() int {
	winningNumbers, elfNumbers := getCards()
	var copies map[int]int = make(map[int]int, len(winningNumbers))

	for i := range winningNumbers {
		copies[i] = 1
	}

	for i := 0; i < len(winningNumbers); i++ {
		winningCard := winningNumbers[i]
		elfCard := elfNumbers[i]

		numWinners := getNumWinningCards(winningCard, elfCard)

		for j := i + 1; j <= i+numWinners; j++ {
			copies[j] += copies[i]
		}
	}

	totalCards := 0

	for _, v := range copies {
		totalCards += v
	}

	return totalCards
}
