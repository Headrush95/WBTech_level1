package main

import (
	"fmt"
	"os"
	"sync"
)

var (
	OUT  = os.Stdout
	DATA = []int{2, 4, 6, 8, 10}
	SUM  = 0
)

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

/*
gettingSquareChan считает квадрат числа и помещает его в канал squares,
после чего отчитывается о выполнении в канал isCompleted
*/
func gettingSquareChan(num int, squares chan<- int, isCompleted chan<- struct{}) {
	squares <- num * num
	isCompleted <- struct{}{}
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

	//через каналы
	{
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

}
