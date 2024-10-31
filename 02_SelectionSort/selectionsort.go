package main

import "fmt"

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}

func selectionSort(array []int) []int {
	size := len(array)
	sorted := make([]int, size)
	for i := 0; i < size; i++ {
		smallest := findSmallest(array)
		sorted[i] = array[smallest]
		// on append tout les élements avant smallest et après en enlevant le smallest
		array = append(array[:smallest], array[smallest+1:]...)
	}
	return sorted
}
func main() {
	s := []int{5, 3, 6, 2, 10}
	fmt.Println(selectionSort(s))
}
