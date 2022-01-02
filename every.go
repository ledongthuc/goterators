package goterators

func Every[T comparable](source []T, conditionFunc func(item T) bool) bool {
	for _, item := range source {
		if !conditionFunc(item) {
			return false
		}
	}
	return true
}
