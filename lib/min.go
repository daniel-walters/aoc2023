package lib

import "math"

func Min(nums ...int) int {
	minNum := math.MaxInt

	for _, v := range nums {
		if v < minNum {
			minNum = v
		}
	}

	return minNum
}
