package goterators

import "golang.org/x/exp/constraints"

// Count returns number of checking item exists in source list
func Count[T constraints.Ordered](source []T, checkedItem T) (result int) {
	for _, item := range source {
		if item == checkedItem {
			result++
		}
	}
	return result
}
