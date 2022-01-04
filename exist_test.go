package goterators

import (
	"fmt"
	"testing"
)

func TestExistInt(t *testing.T) {
	testSource := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	expectedItem1 := 18
	if existed := Exist(testSource, expectedItem1); !existed {
		t.Errorf("Exist item1 = %v, Expected = found , got = not_found", expectedItem1)
	}

	expectedItem2 := 50
	if existed := Exist(testSource, expectedItem2); existed {
		t.Errorf("Exist item1 = %v, Expected = not_found , got = found", expectedItem1)
	}
}

func TestExistString(t *testing.T) {
	testSource := []string{"basis", "pray", "platform", "top", "border", "pole", "kidnap", "eaux", "stream", "lemon", "helpless", "indulge", "highlight", "city", "miner", "photography", "squeeze", "king", "offender", "rugby"}

	expectedItem1 := "king"
	if existed := Exist(testSource, expectedItem1); !existed {
		t.Errorf("Exist item2 = %v, Expected = found , got = not_found", expectedItem1)
	}

	expectedItem2 := "webuild"
	if existed := Exist(testSource, expectedItem2); existed {
		t.Errorf("Exist item2 = %v, Expected = not_found , got = found", expectedItem1)
	}
}

func ExampleExist(t *testing.T) {
	testSource := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println("Exist: ", Exist(testSource, 15))
}
