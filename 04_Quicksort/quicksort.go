package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

func sum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	} else {
		return numbers[0] + sum(numbers[1:])
	}
}
func numberOfItemsInAList(list []int) int {
	if len(list) == 0 {
		return 0
	} else {
		return 1 + numberOfItemsInAList(list[1:])
	}
}

func findMax(list []int) int {
	if len(list) == 1 {
		return list[0]
	}
	if len(list) == 2 {
		if list[0] > list[1] {
			return list[0]
		}
		return list[1]
	}
	subMax := findMax(list[1:])
	if list[0] > subMax {
		return list[0]
	}
	return subMax
}

func binarySearch(list []int, item int) int {
	if len(list) == 1 {
		if list[0] == item {
			return item
		}
		return -1
	}
	low := 0
	high := len(list) - 1
	mid := (low + high) / 2
	if list[mid] >= item {
		return binarySearch(list[0:mid+1], item)
	}
	return binarySearch(list[mid+1:], item)

}

func quicksort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	if len(list) == 2 {
		if list[0] > list[1] {
			return []int{list[1], list[0]}
		}
		return list
	}
	low := 0
	high := len(list) - 1
	mid := (low + high) / 2
	pivot := list[mid]
	lessThan := []int{}
	greaterThan := []int{}
	beforePivot := list[0:mid]
	afterPivot := list[mid+1:]
	pivotExcluded := append(beforePivot, afterPivot...)
	for _, v := range pivotExcluded {
		if v <= pivot {
			lessThan = append(lessThan, v)
		} else {
			greaterThan = append(greaterThan, v)
		}
	}

	lessThan = append(quicksort(lessThan), pivot)
	greaterThan = quicksort(greaterThan)
	return append(
		lessThan, greaterThan...,
	)
}

func main() {
	if sum([]int{2, 4, 6}) != 12 {
		panic("should be 12")
	}
	println(sum([]int{2, 4, 6}))
	nbItems := numberOfItemsInAList([]int{1, 3, 8, 7, 5})
	if nbItems != 5 {
		panic("should be 5")
	}
	max := findMax([]int{1, 3, 8, 7, 5})
	if max != 8 {
		msg := fmt.Errorf("should be 8, was %d", max)
		panic(msg)
	}
	binarySearchResult := binarySearch([]int{2, 3, 8, 12}, 8)
	if binarySearchResult != 8 {
		msg := fmt.Errorf("should be 8, was %d", binarySearchResult)
		panic(msg)

	}
	spew.Dump(quicksort([]int{2, 3, 1, 9, 4, 89, 5}))
	// println(quicksort([]int{2, 3, 1}))
}
