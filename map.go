package goterators

func Map[T1 any, T2 any](source []T1, mappingFunc func(item T1) T2) (output []T2) {
	for _, item := range source {
		output = append(output, mappingFunc(item))
	}
	return output
}
