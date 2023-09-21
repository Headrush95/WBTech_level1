package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в
конкурентной среде. По завершению программа должна выводить итоговое
значение счетчика.
*/

/*
Вариант 2 - с использованием atomic
*/

type counter struct {
	value int64
}

func (c *counter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

func (c *counter) showValue() {
	fmt.Println(c.value)
}

func doIncrement(c *counter, wg *sync.WaitGroup) {
	defer wg.Done()
	c.Increment()
}

func main() {
	start := time.Now()
	wg := &sync.WaitGroup{}
	count := &counter{}
	for i := 0; i < 10000000; i++ {
		wg.Add(1)
		go doIncrement(count, wg)
	}

	wg.Wait()
	count.showValue()
	fmt.Println(time.Since(start).Nanoseconds())
}
