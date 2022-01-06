package goterators

import "constraints"

// CountList returns sub-list counter of input sub-list that want to count from source list.
func CountList[T constraints.Ordered](source []T, checkedItems []T) []int {
	result := make([]int, len(checkedItems))
	for index, item := range checkedItems {
		result[index] = Count(source, item)
	}
	return result
}
