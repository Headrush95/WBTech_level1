package main

import (
	"fmt"
	"log"
)

/*
Удалить i-ый элемент из слайса.
*/

/*
Вариант 1 - через append.
*/
func deleteFromSlice[T any](arr []T, idx int) []T {
	if idx < 0 {
		log.Println("index cannot less then zero")
		return arr
	}
	if idx >= len(arr) {
		log.Println("index out of array length range")
		return arr
	}
	// элементы, начиная с idx+1 и до конца, сдвигаются на 1 позицию
	return append(arr[:idx], arr[idx+1:]...)
}
func main() {
	src := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(src, len(src), cap(src))
	res := deleteFromSlice(src, 2)
	fmt.Println(res, len(res), cap(res))

}
