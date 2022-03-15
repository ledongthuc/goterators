package goterators

import "golang.org/x/exp/constraints"

// Min find smallest value from source list
func Min[T constraints.Ordered](source []T) (result T, err error) {
	if len(source) == 0 {
		return result, ErrNotFound
	}
	result = source[0]
	for _, item := range source {
		if item < result {
			result = item
		}
	}
	return result, nil
}
