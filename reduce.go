package goterators

/*Reduce runs the reducer function on each element of the array. In order, the reduce function passes in the return value from the calculation on the preceding element. The final result of running the reducer across all elements of the array is the return value of the final reducer on the last element.
- list paramerter: source list we want to process.
- initial value paramerter: the previous value that's used in the reducer call of the first element. At this time, previous = initial value, current = first element of the list.
- reducer function paramerter: the function will run on all elements of the source list.*/
func Reduce[T1 any, T2 any](source []T1, initialValue T2, reducer func(previousValue T2, currentValue T1, currentIndex int, list []T1) T2) T2 {
	previousItem := initialValue
	output := initialValue
	for index := range source {
		output = reducer(previousItem, source[index], index, source)
		previousItem = output
	}
	return output
}
