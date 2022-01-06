package goterators

import (
	"fmt"
	"testing"
)

func TestMin(t *testing.T) {
	testSource := []int{20, 17, 9, 21, 18, 3, 11, 5}
	expectedValue := 3
	actutalValue, _ := Min(testSource)

	if actutalValue != expectedValue {
		t.Errorf("Expected = %v , got = %v", expectedValue, actutalValue)
	}
}

func ExampleMin() {
	testSource := []int{20, 17, 9, 21, 18, 3, 11, 5}
	result, _ := Min(testSource)
	fmt.Println("Min: ", result)
}
