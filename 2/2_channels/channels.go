package main

import (
	"fmt"
	"os"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout
*/
var (
	OUT  = os.Stdout
	DATA = []int{2, 4, 6, 8, 10}
)

/*
Вариант 2 - с использованием каналов
*/

/*
gettingSquareDoubleChan считает квадрат числа и помещает его в канал squares,
после чего отчитывается о выполнении в канал isCompleted
*/
func gettingSquareDoubleChan(n int, squares chan<- int, isCompleted chan<- struct{}) {
	squares <- n * n
	isCompleted <- struct{}{}
}

func main() {
	//через каналы
	squares := make(chan int, len(DATA))

	// откидываем писателя в отдельую горутину
	go func() {
		defer close(squares)
		// isCompleted - канал, который служит для гарантии, что все горутиный gettingSquareChan завершили работу
		isCompleted := make(chan struct{}, len(DATA))
		defer close(isCompleted)

		for _, v := range DATA {
			go gettingSquareDoubleChan(v, squares, isCompleted)
		}

		// дожидаемся выполнения всех горутин gettingSquareChan
		for i := 0; i < len(DATA); i++ {
			<-isCompleted
		}
	}()

	// читаем из канала до его закрытия
	for sq := range squares {
		fmt.Fprintf(OUT, "%d ", sq)
	}
}
