package binarysearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	list := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	item := 23

	index := binarySearch(list, item)
	if index != 5 {
		t.Errorf("Unexpected index: %d", index)
	}

	item = 100
	index = binarySearch(list, item)
	if index != -1 {
		t.Errorf("Unexpected index: %d", index)
	}
}
