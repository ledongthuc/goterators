package goterators

func Find[T any](source []T, matchedItemFunc func(item T) bool) (t T, index int, err error) {
	for index, item := range source {
		if matchedItemFunc(item) {
			return item, index, nil
		}
	}
	return t, -1, ErrNotFound
}
