package goterators

import (
	"fmt"
	"testing"
)

func TestReduceSum(t *testing.T) {
	testSource := []int{1, 2, 3, 4, 5}
	expectedValue := 15

	total := Reduce(testSource, 0, func(previous int, current int, index int, list []int) int {
		return previous + current
	})
	if total != expectedValue {
		t.Errorf("Expected = %v, got = %v", expectedValue, total)
	}
}

func TestReduceSimulateMap(t *testing.T) {
	testSource := []int{1, 2, 3, 4, 5}
	expectedItems := []float64{1.0, 2.0, 3.0, 4.0, 5.0}

	items := Reduce(testSource, []float64{}, func(previous []float64, current int, index int, list []int) []float64 {
		return append(previous, float64(current))
	})
	if len(expectedItems) != len(items) {
		t.Fatalf("Expected length = %v, actual length = %v", len(expectedItems), len(items))
	}
	for index := range expectedItems {
		if expectedItems[index] != items[index] {
			t.Errorf("Index %v, expected = %v, actual = %v", index, expectedItems[index], items[index])
		}
	}
}

func ExampleReduce() {
	testSource := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	result := Reduce(testSource, 0, func(previous int, current int, index int, list []int) int {
		return previous + current
	})
	fmt.Println("Reduce: ", result)
}
