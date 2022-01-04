package goterators

import (
	"fmt"
	"testing"
)

func TestFlat(t *testing.T) {
	testSource := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8, 9, 10, 11}, {12, 13, 14, 15, 16, 17, 18, 19}, {20}, {}}
	expectedItems := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	actualItems := Flat(testSource)
	if len(expectedItems) != len(actualItems) {
		t.Fatalf("Expected length = %v, actual length = %v", len(expectedItems), len(actualItems))
	}
	for index := range expectedItems {
		if expectedItems[index] != actualItems[index] {
			t.Errorf("Index %v, expected = %v, actual = %v", index, expectedItems[index], actualItems[index])
		}
	}
}

func TestFlatNested(t *testing.T) {
	testSource := [][][]int{{{1, 2, 3, 4}, {5, 6, 7, 8, 9, 10, 11}, {12, 13, 14, 15, 16, 17, 18, 19}}, {{20}}, {{}}}
	expectedItems := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8, 9, 10, 11}, {12, 13, 14, 15, 16, 17, 18, 19}, {20}, {}}

	actualItems := Flat(testSource)
	if len(expectedItems) != len(actualItems) {
		t.Fatalf("Expected length = %v, actual length = %v", len(expectedItems), len(actualItems))
	}
	for groupIndex := range expectedItems {
		if len(expectedItems[groupIndex]) != len(actualItems[groupIndex]) {
			t.Fatalf("Group index %v, Expected length = %v, actual length = %v", groupIndex, len(expectedItems), len(actualItems))
		}

		for index := range expectedItems[groupIndex] {
			if expectedItems[groupIndex][index] != actualItems[groupIndex][index] {
				t.Errorf("Group index %v, Index %v, expected = %v, actual = %v", groupIndex, index, expectedItems[groupIndex][index], actualItems[groupIndex][index])
			}
		}
	}
}

func ExampleFlat(t *testing.T) {
	testSource := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8, 9, 10, 11}, {12, 13, 14, 15, 16, 17, 18, 19}, {20}, {}}
	items := Flat(testSource)
	fmt.Println("Flat: ", items)
}
