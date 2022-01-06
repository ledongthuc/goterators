package goterators

// Range return max - min
func Range[T Number](source []T) T {
	if len(source) <= 1 {
		var result T
		return result
	}
	max, _ := Max(source)
	min, _ := Min(source)
	return max - min
}
