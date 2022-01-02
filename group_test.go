package goterators

import "testing"

func TestGroup(t *testing.T) {
	type Product struct {
		productType string
		productName string
	}
	testSource := []Product{{
		productType: "food",
		productName: "Coca-laco",
	}, {
		productType: "sport",
		productName: "Old Balance",
	}, {
		productType: "sport",
		productName: "Noke",
	}, {
		productType: "food",
		productName: "Chip chip",
	}}
	expectedItems := [][]Product{
		{
			{
				productType: "food",
				productName: "Coca-laco",
			}, {
				productType: "food",
				productName: "Chip chip",
			},
		},
		{
			{
				productType: "sport",
				productName: "Old Balance",
			}, {
				productType: "sport",
				productName: "Noke",
			},
		},
	}

	actualItems := Group(testSource, func(item Product) string {
		return item.productType
	})
	if len(expectedItems) != len(actualItems) {
		t.Fatalf("Expected length = %v, actual length = %v", len(expectedItems), len(actualItems))
	}
	for groupIndex := range expectedItems {
		if len(expectedItems[groupIndex]) != len(actualItems[groupIndex]) {
			t.Fatalf("Group index %v, expected = %v, actual = %v", groupIndex, expectedItems[groupIndex], actualItems[groupIndex])
		}
		for index := range expectedItems[groupIndex] {
			if expectedItems[groupIndex][index] != actualItems[groupIndex][index] {
				t.Errorf("Group index %v, Index %v, expected = %v, actual = %v", groupIndex, index, expectedItems[groupIndex][index], actualItems[groupIndex][index])
			}
		}
	}
}
