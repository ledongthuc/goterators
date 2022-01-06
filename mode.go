package goterators

import "constraints"

// Mode return a value that appears most often in the source list.
func Mode[T constraints.Ordered](source []T) (T, int) {
	counter := map[T]int{}
	mostOftenCounter := 0
	var mostOftenValue T
	for _, item := range source {
		counter[item]++
		itemCounter := counter[item]
		if itemCounter > mostOftenCounter {
			mostOftenCounter = itemCounter
			mostOftenValue = item
		}
	}
	return mostOftenValue, mostOftenCounter
}
