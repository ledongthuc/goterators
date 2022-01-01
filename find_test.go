package goterators

import "testing"

func TestFindInt(t *testing.T) {
	testSource := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	expectedItem1 := 18
	matchedItem, err := Find(testSource, func(item int) bool {
		return item == expectedItem1
	})
	if err != nil {
		t.Errorf("Find item1 = %v, expected no error, got = %v", expectedItem1, err)
	}
	if matchedItem != expectedItem1 {
		t.Errorf("Find item1 = %v, Expected = %v , got = %v", expectedItem1, expectedItem1, matchedItem)
	}

	expectedItem2 := 50
	matchedItem, err = Find(testSource, func(item int) bool {
		return item == expectedItem2
	})
	if err == nil {
		t.Errorf("Find item2 = %v, expected <not_found>, got = no-error", expectedItem1)
	}
}

func TestFindString(t *testing.T) {
	testSource := []string{"basis", "pray", "platform", "top", "border", "pole", "kidnap", "eaux", "stream", "lemon", "helpless", "indulge", "highlight", "city", "miner", "photography", "squeeze", "king", "offender", "rugby"}

	expectedItem1 := "king"
	matchedItem, err := Find(testSource, func(item string) bool {
		return item == expectedItem1
	})
	if err != nil {
		t.Errorf("Find item1 = %v, expected no error, got = %v", expectedItem1, err)
	}
	if matchedItem != expectedItem1 {
		t.Errorf("Find item1 = %v, Expected = %v , got = %v", expectedItem1, expectedItem1, matchedItem)
	}

	expectedItem2 := "webuild"
	matchedItem, err = Find(testSource, func(item string) bool {
		return item == expectedItem2
	})
	if err == nil {
		t.Errorf("Find item2 = %v, expected <not_found>, got = no-error", expectedItem1)
	}
}
