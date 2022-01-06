package goterators

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	testSource := []int{20, 17, 9, 21, 18, 3, 11, 5}
	expectedValue := 21
	actutalValue, _ := Max(testSource)

	if actutalValue != expectedValue {
		t.Errorf("Expected = %v , got = %v", expectedValue, actutalValue)
	}
}

func ExampleMax() {
	testSource := []int{20, 17, 9, 21, 18, 3, 11, 5}
	result, _ := Max(testSource)
	fmt.Println("Max: ", result)
}
