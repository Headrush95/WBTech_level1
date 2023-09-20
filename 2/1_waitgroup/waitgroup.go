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
Вариант 1 - с использованием sync.WaitGroup
*/

// gettingSquareWG считает квадрат числа и выводит его
func gettingSquareWG(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintf(OUT, "%d ", n*n)
}

func main() {
	// через WaitGroup
	wg := &sync.WaitGroup{}
	for _, v := range DATA {
		wg.Add(1)
		go gettingSquareWG(v, wg)
	}
	wg.Wait()
}
