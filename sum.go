package goterators

// Sum plus all item from source list
func Sum[T Number](source []T) (result T) {
	for _, item := range source {
		result += item
	}
	return result
}
