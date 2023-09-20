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
Вариант 2 - средства стандартной библиотеки sync.Map
*/

func sMapReader(sMap *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		fmt.Println(sMap.Load(i))
	}
}

func sMapWriter(sMap *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		sMap.Store(rand.Intn(1000), i)
	}
}

func main() {
	var sMap sync.Map
	wg := &sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		sMap.Store(i, rand.Intn(i+1))
	}
	// откидываем читателя и писателя в отдельные горутины
	wg.Add(2)
	go sMapReader(&sMap, wg)
	go sMapWriter(&sMap, wg)

	wg.Wait()
}
