package goterators

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	testSource := []int{1, 1, 1, 2, 2, 3, 3, 3, 3, 4}

	findingItem := 3
	expectedValue := 4
	actutalValue := Count(testSource, findingItem)
	if actutalValue != expectedValue {
		t.Errorf("Expected = %v , got = %v", expectedValue, actutalValue)
	}

	findingItem = 5
	expectedValue = 0
	actutalValue = Count(testSource, findingItem)
	if actutalValue != expectedValue {
		t.Errorf("Expected = %v , got = %v", expectedValue, actutalValue)
	}
}

func ExampleCount() {
	testSource := []int{1, 1, 1, 2, 2, 3, 3, 3, 3, 4}
	fmt.Println("Count: ", Count(testSource, 3))
}
