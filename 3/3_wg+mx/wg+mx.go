package main

import (
	"fmt"
	"os"
	"sync"
)

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.
*/

var (
	OUT  = os.Stdout
	DATA = []int{2, 4, 6, 8, 10}
	SUM  = 0
)

/*
Варианты 3.1 и 3.2 - с использованием sync.WaitGroup и sync.Mutex. Различия в области видимости переменной sum
*/

// gettingSquareWGMxGlobal считает квадрат числа и суммирует с глобальной переменной SUM
func gettingSquareWGMxGlobal(num int, wg *sync.WaitGroup, mx *sync.Mutex) {
	defer wg.Done()
	mx.Lock()
	defer mx.Unlock()
	SUM += num * num
}

// gettingSquareWGMxLocal считает квадрат числа и суммирует с локальной переменной sum
func gettingSquareWGMxLocal(num int, sum *int, wg *sync.WaitGroup, mx *sync.Mutex) {
	defer wg.Done()
	mx.Lock()
	defer mx.Unlock()
	*sum += num * num
}

func main() {
	// через WaitGroup, Mutex и локальную переменную
	{
		sum := 0
		wg := &sync.WaitGroup{}
		mx := &sync.Mutex{}
		for _, v := range DATA {
			wg.Add(1)
			go gettingSquareWGMxLocal(v, &sum, wg, mx)
		}
		wg.Wait()
		fmt.Fprintf(OUT, "[WGL] %d\n", sum)
	}

	// через WaitGroup, Mutex и глобальную переменную
	{
		wg := &sync.WaitGroup{}
		mx := &sync.Mutex{}
		for _, v := range DATA {
			wg.Add(1)
			go gettingSquareWGMxGlobal(v, wg, mx)
		}
		wg.Wait()
		fmt.Fprintf(OUT, "[WGG] %d\n", SUM)
	}
}
