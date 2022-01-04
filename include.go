package goterators

// Include check if source list contains all items from the sub-list item.
func Include[T comparable](source, subList []T) bool {
	for _, checkingItem := range subList {
		if existed := Exist(source, checkingItem); !existed {
			return false
		}
	}
	return true
}
