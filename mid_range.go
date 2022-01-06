package goterators

// MidRange return (max + min) / 2
func MidRange[T Number](source []T) float64 {
	if len(source) <= 1 {
		return 0.0
	}
	max, _ := Max(source)
	min, _ := Min(source)
	return float64(max+min) / 2.0
}
