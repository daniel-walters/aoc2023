package main

import (
	"aoc/lib"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type hand struct {
	red   int
	green int
	blue  int
}
type config hand
type game []hand

func newHand(r, g, b int) hand {
	return hand{r, g, b}
}

func (g game) getMinCubes() hand {
	maxR := -1
	maxG := -1
	maxB := -1

	for _, h := range g {
		maxR = max(maxR, h.red)
		maxG = max(maxG, h.green)
		maxB = max(maxB, h.blue)
	}

	return newHand(maxR, maxG, maxB)
}

func (cfg config) handIsPossible(hand hand) bool {
	return hand.red <= cfg.red && hand.green <= cfg.green && hand.blue <= cfg.blue
}

func main() {
	partOneAnswer := partOne(config{12, 13, 14})
	partTwoAnswer := partTwo(config{12, 13, 14})

	lib.LogAnswers(partOneAnswer, partTwoAnswer)
}

func partOne(cfg config) int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := bufio.NewScanner(file)

	sum := 0

	for i := 1; lines.Scan(); i++ {
		line := lines.Text()
		game := getGame(line)
		gamePossible := true

		for _, hand := range game {
			if !cfg.handIsPossible(hand) {
				gamePossible = false
			}
		}

		if gamePossible {
			sum += i
		}
	}

	return sum
}

func getGame(line string) game {
	hands := strings.FieldsFunc(line, func(ch rune) bool {
		return ch == ':' || ch == ';'
	})[1:]

	var game game

	for _, hand := range hands {
		game = append(game, getHand(hand))
	}

	return game
}

func getHand(hand string) hand {
	cubes := strings.Split(hand, ",")
	out := newHand(0, 0, 0)

	for _, cube := range cubes {
		cube = strings.Trim(cube, " ")

		parts := strings.Split(cube, " ")

		num, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("Failed to parse to int: %s", parts[0])
		}

		switch parts[1] {
		case "red":
			out.red = num
		case "green":
			out.green = num
		case "blue":
			out.blue = num
		}
	}

	return out
}

func partTwo(cfg config) int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := bufio.NewScanner(file)

	sum := 0

	for lines.Scan() {
		line := lines.Text()
		game := getGame(line)
		minCubeHand := game.getMinCubes()

		product := minCubeHand.red * minCubeHand.blue * minCubeHand.green

		sum += product
	}

	return sum
}
