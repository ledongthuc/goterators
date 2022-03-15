package goterators

import (
	"fmt"
	"testing"
)

func TestForEach(t *testing.T) {
	testForEach([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, t)
	testForEach([]string{"basis", "pray", "platform", "top", "border", "pole", "kidnap", "eaux", "stream", "lemon", "helpless", "indulge", "highlight", "city", "miner", "photography", "squeeze", "king", "offender", "rugby"}, t)
}

func testForEach[K comparable](source []K, t *testing.T) {
	indexCounter := 0
	ForEach(source, func(item K) {
		if indexCounter >= len(source) {
			t.Fatalf("Expected max index = %v, but got index = %v", len(source), indexCounter)
		}
		if source[indexCounter] != item {
			t.Errorf("Index %v - Expected: %+v, but got: %+v", indexCounter, source[indexCounter], item)
		}
		indexCounter++
	})

	if indexCounter != len(source) {
		t.Fatalf("Expected number of items = %v, but got =  %v", len(source), indexCounter)
	}
}

func ExampleForEach() {
	testSource := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	ForEach(testSource, func(item int) {
		fmt.Println("ForEach: ", item)
	})
}
