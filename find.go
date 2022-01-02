package goterators

// Find returns the first element and its index in the list that meets the functional condition. If no element meet the condition function, return the error "Not Found".
func Find[T any](source []T, matchedItemFunc func(item T) bool) (t T, index int, err error) {
	for index, item := range source {
		if matchedItemFunc(item) {
			return item, index, nil
		}
	}
	return t, -1, ErrNotFound
}
