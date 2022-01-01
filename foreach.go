package goterators

func ForEach[T any](source []T, handler func(item T)) {
	for _, item := range source {
		handler(item)
	}
}
