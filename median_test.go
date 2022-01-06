package goterators

import (
	"fmt"
	"testing"
)

func TestMedian(t *testing.T) {
	testSource := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	expectedValue := 11
	expectedIndex := 10
	actutalValue, actualIndex, _ := Median(testSource)
	if actutalValue != expectedValue {
		t.Errorf("Expected = %v , got = %v", expectedValue, actutalValue)
	}
	if actualIndex != expectedIndex {
		t.Errorf("Expected = %v , got = %v", expectedIndex, actualIndex)
	}

	testSource = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	expectedValue = 10
	expectedIndex = 9
	actutalValue, actualIndex, _ = Median(testSource)
	if actutalValue != expectedValue {
		t.Errorf("Expected = %v , got = %v", expectedValue, actutalValue)
	}
	if actualIndex != expectedIndex {
		t.Errorf("Expected = %v , got = %v", expectedIndex, actualIndex)
	}
}

func ExampleMedian() {
	testSource := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	median, index, _ := Median(testSource)
	fmt.Println("Median: ", median, ", with index: ", index)
}
