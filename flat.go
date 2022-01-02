package goterators

// Flat returns a new array with all sub-array elements concatenated with 1 level depth.
func Flat[T any](source [][]T) []T {
	outputSize := 0
	for _, group := range source {
		outputSize += len(group)
	}
	output := make([]T, 0, outputSize)

	for _, group := range source {
		output = append(output, group...)
	}
	return output
}
