package main

import (
	"aoc/lib"
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	partOneAnswer := solution(false)
	partTwoAnswer := solution(true)

	lib.LogAnswers(partOneAnswer, partTwoAnswer)
}

func getFirstNumber(str string) int {
	for _, char := range str {
		if unicode.IsDigit(char) {
			return int(char - '0')
		}
	}

	return -1
}

func replaceNumberWords(str string) string {
	chars := []rune(str)
	clonedString := strings.Clone(str)

	for k, v := range numberConv {
		for i := strings.Index(clonedString, k); i != -1; i = strings.Index(clonedString, k) {
			chars[i+1] = rune('0' + v)
			clonedString = string(chars)
		}
	}

	return clonedString
}

func reverseString(str string) string {
	chars := []rune(str)

	slices.Reverse(chars)

	return string(chars)
}

func combineNumbers(a, b int) (int, error) {
	combinedNumber := fmt.Sprintf("%d%d", a, b)

	return strconv.Atoi(combinedNumber)
}

var numberConv = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func solution(parseWords bool) int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := bufio.NewScanner(file)

	sum := 0

	for lines.Scan() {
		line := lines.Text()

		if parseWords {
			line = replaceNumberWords(line)
		}

		firstNumber := getFirstNumber(line)
		lastNumber := getFirstNumber(reverseString(line))

		combinedNumber, err := combineNumbers(firstNumber, lastNumber)
		if err != nil {
			log.Fatalf("Failed to combine numbers %d & %d", firstNumber, lastNumber)
		}

		sum += combinedNumber
	}

	return sum
}
