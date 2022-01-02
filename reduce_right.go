package goterators

/* ReduceRight runs the reducer function on each element of the array, from last to the first element. In order, the reduce function passes in the return value from the calculation on the preceding element. The final result of running the reducer across all elements of the array is the return value of the final reducer on the first element.
- list parameter: source list we want to process.
- initial value parameter: the previous value that's used in the reducer call of the last element. At this time, previous = initial value, current = last element of list.
- reducer function parameter: the function will run on all elements of the source list.*/
func ReduceRight[T1 any, T2 any](source []T1, initialValue T2, reducer func(previousValue T2, currentValue T1, currentIndex int, list []T1) T2) T2 {
	previousItem := initialValue
	output := initialValue
	for index := len(source) - 1; index >= 0; index-- {
		output = reducer(previousItem, source[index], index, source)
		previousItem = output
	}
	return output
}
