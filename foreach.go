package goterators

// ForEach does the same `for` in Go. Just another option to loop through items in a list.
func ForEach[T any](source []T, handler func(item T)) {
	for _, item := range source {
		handler(item)
	}
}
