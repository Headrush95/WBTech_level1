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
Вариант 1 - с использованием двух каналов: с данными и канал завершения
*/

/*
gettingSquareChan считает квадрат числа и помещает его в канал squares,
после чего отчитывается о выполнении в канал isCompleted
*/
func gettingSquareChan(num int, squares chan<- int, isCompleted chan<- struct{}) {
	squares <- num * num
	isCompleted <- struct{}{}
}

func main() {
	//через каналы (с каналом завершения)
	squares := make(chan int, len(DATA))
	sum := 0

	// откидываем писателя в отдельую горутину
	go func() {
		defer close(squares)
		// isCompleted - канал, который служит для гарантии, что все горутиный gettingSquareChan завершили работу
		isCompleted := make(chan struct{}, len(DATA))
		defer close(isCompleted)

		for _, v := range DATA {
			go gettingSquareChan(v, squares, isCompleted)
		}

		// дожидаемся выполнения всех горутин gettingSquareChan
		for i := 0; i < len(DATA); i++ {
			<-isCompleted
		}

	}()

	for sq := range squares {
		sum += sq
	}

	fmt.Fprintf(OUT, "[Chan] %d\n", sum)
}
