package main

import (
	"fmt"
	"os"
)

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.
*/

var (
	OUT  = os.Stdout
	DATA = []int{2, 4, 6, 8, 10}
)

/*
Вариант 2 - используется один канал
*/

func getSquaresSingleChan(num int, sq chan<- int) {
	sq <- num * num
}

func main() {
	squares := make(chan int, len(DATA))
	defer close(squares)
	sum := 0

	go func() {
		for _, num := range DATA {
			go getSquaresSingleChan(num, squares)
		}
	}()

	for i := 0; i < len(DATA); i++ {
		sum += <-squares
	}
	fmt.Fprintf(OUT, "[SChan] %d", sum)
}
