package goterators

import (
	"fmt"
	"testing"
)

func TestMode(t *testing.T) {
	testSource := []int{1, 1, 1, 2, 2, 2, 2, 3, 3, 4}
	expectedValue := 2
	expectedCounter := 4

	actualValue, actualCounter := Mode(testSource)
	if actualValue != expectedValue {
		t.Errorf("Expected = %v , got = %v", expectedValue, actualValue)
	}
	if actualCounter != expectedCounter {
		t.Errorf("Expected = %v , got = %v", expectedCounter, actualCounter)
	}

}

func ExampleMode() {
	testSource := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	mostOftenValue, counter := Mode(testSource)
	fmt.Println("Mode: ", mostOftenValue, counter)
}
