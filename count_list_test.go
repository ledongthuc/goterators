package goterators

import (
	"fmt"
	"testing"
)

func TestCountList(t *testing.T) {
	testSource := []int{1, 1, 1, 2, 2, 3, 3, 3, 3, 4}

	countingItems := []int{1, 1, 2, 3, 5}
	expectedValue := []int{3, 3, 2, 4, 0}
	actualValue := CountList(testSource, countingItems)
	if len(actualValue) != len(expectedValue) {
		t.Fatalf("Expected length = %v , got length = %v", len(expectedValue), len(actualValue))
	}
	for index := range expectedValue {
		if expectedValue[index] != actualValue[index] {
			t.Errorf("Expected = %v , got = %v", expectedValue[index], actualValue[index])
		}
	}
}

func ExampleCountList() {
	testSource := []int{1, 1, 1, 2, 2, 3, 3, 3, 3, 4}
	fmt.Println("CountList: ", CountList(testSource, []int{1, 1, 2, 3, 5}))
}
