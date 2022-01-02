package goterators

func Exist[T comparable](source []T, checkingItem T) bool {
	_, _, err := Find[T](source, func(item T) bool {
		return item == checkingItem
	})
	return err == nil
}
