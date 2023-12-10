package main

import (
	"aoc/lib"
	"bufio"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	partOneAnswer := partOne()
	partTwoAnswer := partTwo()

	lib.LogAnswers(partOneAnswer, partTwoAnswer)
}

func getEngine() [][]rune {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := bufio.NewScanner(file)

	var engine [][]rune

	for lines.Scan() {
		line := lines.Text()
		engine = append(engine, []rune(line))
	}

	return engine
}

func isDigit(char rune) bool {
	return unicode.IsDigit(char)
}

func isSymbol(char rune) bool {
	return char != '.' && !isDigit(char)
}

func isGear(char rune) bool {
	return char == '*'
}

var dirs [][]int = [][]int{
	{0, -1},
	{1, -1},
	{1, 0},
	{1, 1},
	{0, 1},
	{-1, 1},
	{-1, 0},
	{-1, -1},
}

func canAccess(row, col, x, y, maxX, maxY int) bool {
	return row+y < maxY && row+y >= 0 && col+x < maxX && col+x >= 0
}

func isAdjacentToSymbol(row, col int, engine [][]rune) bool {
	for _, dir := range dirs {
		x := dir[0]
		y := dir[1]

		if canAccess(row, col, x, y, len(engine[0]), len(engine)) {
			if isSymbol(engine[row+y][col+x]) {
				return true
			}
		}
	}

	return false
}

func partOne() int {
	engine := getEngine()
	sum := 0

	for i := 0; i < len(engine); i++ {
		row := engine[i]

		for j := 0; j < len(row); j++ {
			char := row[j]

			if isSymbol(char) {
				if nums := isAdjacentToNumber(i, j, engine); len(nums) > 0 {
					for _, num := range nums {
						sum += num
					}
				}
			}

		}
	}

	return sum
}

func getDigit(row, col int, engine [][]rune) int {
	targetCol := col
	numCollector := string(engine[row][col])
	engine[row][targetCol] = '.'
	targetCol--

	for targetCol >= 0 {
		char := engine[row][targetCol]
		if isDigit(char) {
			numCollector = string(char) + numCollector
			engine[row][targetCol] = '.'
			targetCol--
			continue
		}

		break
	}

	targetCol = col + 1

	for targetCol < len(engine[0]) {
		char := engine[row][targetCol]
		if isDigit(char) {
			numCollector = numCollector + string(char)
			engine[row][targetCol] = '.'
			targetCol++
			continue
		}

		break
	}

	num, err := strconv.Atoi(numCollector)
	if err != nil {
		log.Fatalf("Failed to convert to int: %s", numCollector)
	}

	return num
}

func isAdjacentToNumber(row, col int, engine [][]rune) []int {
	nums := []int{}

	for _, dir := range dirs {
		x := dir[0]
		y := dir[1]

		if canAccess(row, col, x, y, len(engine[0]), len(engine)) {
			char := engine[row+y][col+x]

			if isDigit(char) {
				num := getDigit(row+y, col+x, engine)

				nums = append(nums, num)
			}
		}
	}

	return nums
}

func partTwo() int {
	engine := getEngine()
	sum := 0

	for i := 0; i < len(engine); i++ {
		row := engine[i]

		for j := 0; j < len(row); j++ {
			char := row[j]

			if isGear(char) {
				if nums := isAdjacentToNumber(i, j, engine); len(nums) == 2 {
					product := nums[0] * nums[1]
					sum += product
				}
			}

		}
	}

	return sum
}
