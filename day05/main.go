package main

import (
	"aoc/lib"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	partOneAnswer := partOne()
	partTwoAnswer := partOne()

	lib.LogAnswers(partOneAnswer, partTwoAnswer)
}

func getSeeds(line string) []int {
	seeds := []int{}

	seedNums := strings.Split(line, " ")[1:]

	for _, v := range seedNums {
		asInt, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Failed to parse string to int: %s", v)
		}

		seeds = append(seeds, asInt)
	}

	return seeds
}

func getMaps(mapsInput string) [][][]int {
	maps := [][][]int{}
	mapsStrings := strings.Split(mapsInput, "\n\n")

	for _, lines := range mapsStrings {

		mapsToAdd := [][]int{}

		curMaps := strings.Split(lines, "\n")

		for _, curMap := range curMaps[1:] {
			nums := strings.Fields(curMap)
			curMap := []int{}

			for _, num := range nums {
				asInt, err := strconv.Atoi(num)
				if err != nil {
					log.Fatalf("Failed to parse string to int: %s", num)
				}

				curMap = append(curMap, asInt)
			}

			mapsToAdd = append(mapsToAdd, curMap)
		}

		maps = append(maps, mapsToAdd)
	}

	return maps
}

func partOne() int {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	seedsLine, mapLines, _ := strings.Cut(string(file), "\n\n")
	seeds := getSeeds(seedsLine)

	fmt.Printf("seeds:\n%v\n\n", seeds)

	maps := getMaps(mapLines)

	for i, category := range maps {
		fmt.Printf("Category %d\n", i+1)
		for _, m := range category {
			fmt.Printf("map: %v\n", m)
		}
	}

	return 0
}

func partTwo() int {
	return 0
}
