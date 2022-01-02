package goterators

// Some checks at least one element in the list meet the condition; return true, or return false if all elements don't meet the condition.
func Some[T comparable](source []T, conditionFunc func(item T) bool) bool {
	for _, item := range source {
		if conditionFunc(item) {
			return true
		}
	}
	return false
}
