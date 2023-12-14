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
	partTwoAnswer := partTwo()

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

func getSeedsRange(line string) []seedRange {
	seeds := []seedRange{}

	seedsNums := strings.Split(line, " ")[1:]

	for i := 0; i < len(seedsNums)/2; i += 2 {
		start := seedsNums[i]
		seedRange := seedsNums[i+1]

		seedRangeNum, err := strconv.Atoi(seedRange)
		if err != nil {
			log.Fatalf("Could not convert string to num: %s", seedRange)
		}

		startNum, err := strconv.Atoi(start)
		if err != nil {
			log.Fatalf("Could not convert string to num: %s", seedRange)
		}

		seeds = append(seeds, newSeedRange(startNum, seedRangeNum))
	}

	return seeds
}

type seedRange struct {
	start int
	steps int
}

func newSeedRange(start, steps int) seedRange {
	return seedRange{start, steps}
}

type seedMap struct {
	destStart   int
	sourceStart int
	steps       int
}

func newSeedMap(nums []string) seedMap {
	if len(nums) != 3 {
		log.Panicf("Could not create seedMap without 3 args: %v", nums)
	}

	asInts, err := lib.StringsToInts(nums)
	if err != nil {
		log.Fatalln(err)
	}

	return seedMap{asInts[0], asInts[1], asInts[2]}
}

func (s seedMap) String() string {
	return fmt.Sprintf("dest: %d, source: %d, range: %d", s.destStart, s.sourceStart, s.steps)
}

func getDestForSource(source int, maps []seedMap) int {
	for _, m := range maps {
		if source >= m.sourceStart && source < m.sourceStart+m.steps {
			return m.destStart + (source - m.sourceStart)
		}
	}

	return source
}

func getMaps(mapsInput string) [][]seedMap {
	maps := [][]seedMap{}
	mapsStrings := strings.Split(mapsInput, "\n\n")

	for _, lines := range mapsStrings {

		mapsToAdd := []seedMap{}

		curMaps := strings.Split(lines, "\n")

		for _, curMap := range curMaps[1:] {
			if strings.TrimSpace(curMap) == "" {
				continue
			}

			nums := strings.Fields(curMap)

			mapsToAdd = append(mapsToAdd, newSeedMap(nums))
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
	maps := getMaps(mapLines)

	for _, mapList := range maps {
		for i := 0; i < len(seeds); i++ {
			seeds[i] = getDestForSource(seeds[i], mapList)
		}
	}

	return lib.Min(seeds...)
}

// iterate through each seed range
func partTwo() int {
	// file, err := os.ReadFile("input.txt")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// seedsLine, mapLines, _ := strings.Cut(string(file), "\n\n")
	// seeds := getSeedsRange(seedsLine)
	// fmt.Printf("%v", seeds)
	// maps := getMaps(mapLines)

	// for _, mapList := range maps {
	// 	for i := 0; i < len(seeds); i++ {
	// 		seeds[i] = getDestForSource(seeds[i], mapList)
	// 	}
	// }
	//
	// return lib.Min(seeds...)
	return 0
}
