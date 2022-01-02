package goterators

func Group[T any, G comparable](source []T, buildGroup func(item T) G) [][]T {
	output := map[G][]T{}
	for index, item := range source {
		group := buildGroup(item)
		output[group] = append(output[group], source[index])
	}
	return output
}
