package goterators

import (
	"fmt"
	"testing"
)

func TestAverage(t *testing.T) {
	testSource := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	expectedValue := 10.5

	actualValue := Average(testSource)
	if actualValue != expectedValue {
		t.Errorf("Expected = %v , got = %v", expectedValue, actualValue)
	}

	actualValue = Mean(testSource)
	if actualValue != expectedValue {
		t.Errorf("Expected = %v , got = %v", expectedValue, actualValue)
	}
}

func ExampleAverage() {
	testSource := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println("Average: ", Average(testSource))
	fmt.Println("Mean: ", Mean(testSource))
}
