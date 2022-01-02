package goterators

// Filter return items that pass the filter function.
func Filter[T any](source []T, filteredFunc func(item T) bool) (output []T) {
	for _, item := range source {
		if filteredFunc(item) {
			output = append(output, item)
		}
	}
	return output
}
