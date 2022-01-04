package goterators

import (
	"fmt"
	"testing"
)

func TestIncludeSome(t *testing.T) {
	testSource := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	subList := []int{8, 15, 19}
	expectedResult := true
	result := IncludeSome(testSource, subList)
	if result != expectedResult {
		t.Errorf("Expected = %v , got = %v", expectedResult, result)
	}

	subList = []int{8, 15, 21}
	expectedResult = true
	result = IncludeSome(testSource, subList)
	if result != expectedResult {
		t.Errorf("Expected = %v , got = %v", expectedResult, result)
	}

	subList = []int{21, 22, 23}
	expectedResult = false
	result = IncludeSome(testSource, subList)
	if result != expectedResult {
		t.Errorf("Expected = %v , got = %v", expectedResult, result)
	}
}

func ExampleIncludeSome() {
	testSource := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	subList := []int{8, 15, 19}
	result := IncludeSome(testSource, subList)
	fmt.Println("IncludeSome: ", result)
}
