package goterators

// IncludeSome check if source list contains any items from the sub-list item.
func IncludeSome[T comparable](source, subList []T) bool {
	for _, item := range source {
		for _, checkingItem := range subList {
			if item == checkingItem {
				return true
			}
		}
	}
	return false
}
