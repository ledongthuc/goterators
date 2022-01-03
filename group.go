package goterators

// Group groups elements into the nested level. Use a build-group function to define group type.
func Group[T any, G comparable](source []T, buildGroup func(item T) G) [][]T {
	groupOrder := make([]G, 0, len(source))
	m := map[G][]T{}
	for index, item := range source {
		group := buildGroup(item)
		_, existed := m[group]
		if !existed {
			groupOrder = append(groupOrder, group)
		}
		m[group] = append(m[group], source[index])
	}

	output := make([][]T, 0, len(m))
	for _, key := range groupOrder {
		output = append(output, m[key])
	}
	return output
}
