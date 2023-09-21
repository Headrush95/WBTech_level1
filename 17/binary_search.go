package main

import (
	"cmp"
	"fmt"
	"slices"
)

/*
Реализовать бинарный поиск встроенными методами языка.
*/

// binarySearch осуществляет поиск элемента в отсортированном слайсе
func binarySearch[T cmp.Ordered](arr []T, target T) (int, bool) {
	minimum := 0
	maximum := len(arr) - 1
	for minimum <= maximum {
		mid := (minimum + maximum) / 2

		if arr[mid] > target {
			maximum = mid - 1
			continue
		} else if arr[mid] < target {
			minimum = mid + 1
		} else {
			return mid, true
		}
	}
	return -1, false
}

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	idx, isFound := slices.BinarySearch(input, 8)
	fmt.Printf("Number %d, isFound %v, index %d\n", 8, isFound, idx)
	idx, isFound = binarySearch(input, 8)
	fmt.Printf("Number %d, isFound %v, index %d\n", 8, isFound, idx)
}
