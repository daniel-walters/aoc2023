package lib

func Filter[T any](slice []T, filterFn func(item T) bool) []T {
	var out []T = make([]T, 0, len(slice))

	for _, v := range slice {
		if filterFn(v) {
			out = append(out, v)
		}
	}

	return out
}
