package main

import (
	"fmt"
	"log"
)

/*
Удалить i-ый элемент из слайса.
*/

/*
Вариант 2 - через перенос последнего элемента на место удоляемого и сокращение длины слайса на 1.
Профитно на очень больших слайсах, так как выполняется за константное время, но меняется исходный порядок элементов.
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
	arr[idx] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}

func main() {
	src := []int{1, 2, 3, 4, 5, 6, 7, 8}
	res := deleteFromSlice(src, 2)
	fmt.Println(res, len(res), cap(res))
}
