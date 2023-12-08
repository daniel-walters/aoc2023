package lib

import "fmt"

const green = "\u001b[32m"
const reset = "\u001b[0m"

func LogAnswers(partOne, partTwo any) {
	fmt.Printf("%sPart one:%s %v\n", green, reset, partOne)
	fmt.Printf("%sPart two:%s %v\n", green, reset, partTwo)
}
