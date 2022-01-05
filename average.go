package goterators

// Average sum of all the source list divided by the total number of source list
func Average[T Number](source []T) float64 {
	var sum T
	for _, item := range source {
		sum += item
	}
	return float64(sum) / float64(len(source))
}

func Mean[T Number](source []T) float64 {
	return Average(source)
}
