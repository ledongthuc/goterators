package goterators

// Median return a value in the middle of an ordered source list. If number of item in source is even, return right item. Make sure source list are sorted
func Median[T any](source []T) (T, int, error) {
	if len(source) == 0 {
		var result T
		return result, 0, ErrNotFound
	}
	if len(source) == 1 {
		return source[0], 0, nil
	}
	midNumber := len(source) / 2
	return source[midNumber], midNumber, nil
}
