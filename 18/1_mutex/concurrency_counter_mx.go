package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в
конкурентной среде. По завершению программа должна выводить итоговое
значение счетчика.
*/

/*
Вариант 1 - с использованием mutex
*/

type counter struct {
	sync.Mutex
	value int
}

func (c *counter) increment() {
	c.Lock()
	defer c.Unlock()
	c.value++
}

func (c *counter) showValue() {
	fmt.Println(c.value)
}

func doIncrement(c *counter, wg *sync.WaitGroup) {
	defer wg.Done()
	c.increment()
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
	/*
		atomics 	2950327300 nsec
		mutex 		3056598300 nsec
	*/
}
