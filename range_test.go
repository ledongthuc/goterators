package goterators

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {
	testSource := []int{20, 17, 9, 21, 18, 3, 11, 5}
	expectedValue := 18
	actutalValue := Range(testSource)

	if actutalValue != expectedValue {
		t.Errorf("Expected = %v , got = %v", expectedValue, actutalValue)
	}
}

func ExampleRange() {
	testSource := []int{20, 17, 9, 21, 18, 3, 11, 5}
	fmt.Println("Range: ", Range(testSource))
}
