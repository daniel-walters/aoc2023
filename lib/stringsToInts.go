package lib

import (
	"fmt"
	"strconv"
)

func StringsToInts(in []string) ([]int, error) {
	var out []int = make([]int, 0, len(in))

	for _, v := range in {
		asInt, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse string: %s", v)
		}

		out = append(out, asInt)
	}

	return out, nil
}
