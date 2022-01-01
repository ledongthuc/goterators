package goterators

func Find[T any](source []T, matchedItemFunc func(item T) bool) (t T, err error) {
	for _, item := range source {
		if matchedItemFunc(item) {
			return item, nil
		}
	}
	return t, ErrNotFound
}
