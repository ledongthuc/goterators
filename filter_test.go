package goterators

import (
	"fmt"
	"strings"
	"testing"
)

func TestFilterInt(t *testing.T) {
	testSource := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	expectedItems := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	filteredItems := Filter(testSource, func(item int) bool {
		return item%2 == 0
	})
	if len(expectedItems) != len(filteredItems) {
		t.Fatalf("Expected length = %v, actual length = %v", len(expectedItems), len(filteredItems))
	}
	for index := range expectedItems {
		if expectedItems[index] != filteredItems[index] {
			t.Errorf("Index %v, expected = %v, actual = %v", index, expectedItems[index], filteredItems[index])
		}
	}
}

func TestFilterString(t *testing.T) {
	testSource := []string{"basis", "pray", "platform", "top", "border", "pole", "kidnap", "eaux", "stream", "lemon", "helpless", "indulge", "highlight", "city", "miner", "photography", "squeeze", "king", "offender", "rugby"}
	expectedItems := []string{"platform", "top", "border", "pole", "lemon", "photography", "offender"}

	filteredItems := Filter(testSource, func(item string) bool {
		return strings.Contains(item, "o")
	})
	if len(expectedItems) != len(filteredItems) {
		t.Fatalf("Expected length = %v, actual length = %v", len(expectedItems), len(filteredItems))
	}
	for index := range expectedItems {
		if expectedItems[index] != filteredItems[index] {
			t.Errorf("Index %v, expected = %v, actual = %v", index, expectedItems[index], filteredItems[index])
		}
	}
}

func ExampleFilter() {
	testSource := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	filteredItems := Filter(testSource, func(item int) bool {
		return item%2 == 0
	})
	fmt.Println("Filtered: ", filteredItems)
}
