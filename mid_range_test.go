package goterators

import (
	"fmt"
	"testing"
)

func TestMidRange(t *testing.T) {
	testSource := []int{20, 17, 9, 21, 18, 3, 11, 5}
	expectedValue := 12.0
	actutalValue := MidRange(testSource)

	if actutalValue != expectedValue {
		t.Errorf("Expected = %v , got = %v", expectedValue, actutalValue)
	}
}

func ExampleMidRange() {
	testSource := []int{20, 17, 9, 21, 18, 3, 11, 5}
	fmt.Println("MidRange: ", MidRange(testSource))
}
