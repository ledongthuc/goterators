package goterators

// Map function converts items in the list to the output list.
func Map[T1 any, T2 any](source []T1, mappingFunc func(item T1) T2) (output []T2) {
	for _, item := range source {
		output = append(output, mappingFunc(item))
	}
	return output
}
