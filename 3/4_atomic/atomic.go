package main

import (
	"fmt"
	"os"
	"sync"
	"sync/atomic"
)

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.
*/

var (
	OUT  = os.Stdout
	DATA = []int64{2, 4, 6, 8, 10}
)

/*
Вариант 4 - с использованием атомиков
*/

func main() {
	var sum int64
	wg := &sync.WaitGroup{}
	for _, num := range DATA {
		wg.Add(1)
		go func(num int64) {
			defer wg.Done()
			atomic.AddInt64(&sum, num*num)
		}(num)
	}
	wg.Wait()
	fmt.Fprintf(OUT, "[Atomics] %d", sum)
}
