package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/
/*
Вариант 1 - своя структура со встроенным sync.RWMutex
*/
type customMap struct {
	values map[int]int
	*sync.RWMutex
}

func newCustomMap(m map[int]int) customMap {
	return customMap{values: m, RWMutex: &sync.RWMutex{}}
}

// Add не только добавляет новые значения в customMap, но и меняет уже существующие
func (c *customMap) Add(key, value int) {
	c.Lock()
	defer c.Unlock()
	c.values[key] = value
}

func (c *customMap) Get(key int) int {
	c.RLock()
	defer c.RUnlock()
	return c.values[key]
}

func (c *customMap) Delete(key int) {
	c.Lock()
	defer c.Unlock()
	delete(c.values, key)
}

func CustomMapWriter(cm customMap, wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		cm.Add(rand.Intn(1000), rand.Intn(1000))
	}
	wg.Done()
}

func CustomMapReader(cm customMap, wg *sync.WaitGroup) {
	for i := 1; i < 10; i++ {
		fmt.Println(cm.Get(i))
	}
	wg.Done()
}

func main() {
	cMap := newCustomMap(make(map[int]int, 10))
	// предварительно заполняем карту (чтобы в некоторых случаях не читать пустые значения)
	for i := 0; i < 1000; i++ {
		cMap.Add(i, rand.Intn(1000))
	}
	wg := &sync.WaitGroup{}
	// откидываем в горутины писателя и читателя
	wg.Add(2)
	go CustomMapWriter(cMap, wg)
	go CustomMapReader(cMap, wg)
	wg.Wait()
}
