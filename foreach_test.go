package goterators

import "testing"

func TestForEachInt(t *testing.T) {
	testSource := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	indexCounter := 0
	ForEach(testSource, func(item int) {
		if indexCounter >= len(testSource) {
			t.Fatalf("Expected max index = %v, but got index = %v", len(testSource), indexCounter)
		}
		if testSource[indexCounter] != item {
			t.Errorf("Index %v - Expected: %+v, but got: %+v", indexCounter, testSource[indexCounter], item)
		}
		indexCounter++
	})

	if indexCounter != len(testSource) {
		t.Fatalf("Expected number of items = %v, but got =  %v", len(testSource), indexCounter)
	}
}

func TestForEachString(t *testing.T) {
	testSource := []string{"basis", "pray", "platform", "top", "border", "pole", "kidnap", "eaux", "stream", "lemon", "helpless", "indulge", "highlight", "city", "miner", "photography", "squeeze", "king", "offender", "rugby"}

	indexCounter := 0
	ForEach(testSource, func(item string) {
		if indexCounter >= len(testSource) {
			t.Fatalf("Expected max index = %v, but got index = %v", len(testSource), indexCounter)
		}
		if testSource[indexCounter] != item {
			t.Errorf("Index %v - Expected: %+v, but got: %+v", indexCounter, testSource[indexCounter], item)
		}
		indexCounter++
	})

	if indexCounter != len(testSource) {
		t.Fatalf("Expected number of items = %v, but got =  %v", len(testSource), indexCounter)
	}
}
