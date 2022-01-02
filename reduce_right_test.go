package goterators

import (
	"testing"
)

func TestReduceRight(t *testing.T) {
	testSource := []string{"1", "2", "3", "4"}
	expectedItems := []string{"4", "3", "2", "1"}

	actualItems := ReduceRight(testSource, []string{}, func(previous []string, current string, index int, list []string) []string {
		return append(previous, current)
	})
	if len(expectedItems) != len(actualItems) {
		t.Fatalf("Expected length = %v, actual length = %v", len(expectedItems), len(actualItems))
	}
	for index := range expectedItems {
		if expectedItems[index] != actualItems[index] {
			t.Errorf("Index %v, expected = %v, actual = %v", index, expectedItems[index], actualItems[index])
		}
	}
}
