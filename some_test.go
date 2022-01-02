package goterators

import "testing"

func TestSomeValid(t *testing.T) {
	testSource := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	expectedValid := true
	everyValid := Some(testSource, func(item int) bool { return item%2 == 0 })

	if everyValid != expectedValid {
		t.Errorf("Expected = %v , got = %v", expectedValid, everyValid)
	}
}

func TestSomeNotValid(t *testing.T) {
	testSource := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	expectedValid := false
	everyValid := Every(testSource, func(item int) bool { return item > 20 })

	if everyValid != expectedValid {
		t.Errorf("Expected = %v , got = %v", expectedValid, everyValid)
	}
}
