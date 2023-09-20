package main

import (
	"fmt"
	"os"
	"sync"
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
Вариант 3 - с использованием одного канала
*/

/*
gettingSquareSingleChan считает квадрат числа и помещает его в канал squares
*/
func gettingSquareSingleChan(n int, squares chan<- int, wg *sync.WaitGroup) {
	squares <- n * n
	wg.Done()
}

func main() {
	//через канал и sync.WaitGroup
	squares := make(chan int, len(DATA))

	// откидываем писателя в отдельую горутину
	go func() {
		defer close(squares)

		wg := &sync.WaitGroup{}
		for _, v := range DATA {
			wg.Add(1)
			go gettingSquareSingleChan(v, squares, wg)
		}
		wg.Wait()
	}()

	// отличие от второго варианта - тут при чтении мы опираемся на кол-во элементов,
	// тогда как во втором варианте - на закрытие канала squares
	for i := 0; i < len(DATA); i++ {
		fmt.Fprintf(OUT, "%d ", <-squares)
	}

}
