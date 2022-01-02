package goterators

func ReduceRight[T1 any, T2 any](source []T1, initialValue T2, reducer func(previousValue T2, currentValue T1, currentIndex int, list []T1) T2) T2 {
	previousItem := initialValue
	output := initialValue
	for index := len(source) - 1; index >= 0; index-- {
		output = reducer(previousItem, source[index], index, source)
		previousItem = output
	}
	return output
}
