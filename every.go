package goterators

// Every checks all elements in the list with condition function. If it's yes return true; otherwise, return false.
func Every[T comparable](source []T, conditionFunc func(item T) bool) bool {
	for _, item := range source {
		if !conditionFunc(item) {
			return false
		}
	}
	return true
}
