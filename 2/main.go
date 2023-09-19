package main

import (
	"fmt"
	"os"
	"sync"
)

var (
	OUT  = os.Stdout
	DATA = []int{2, 4, 6, 8, 10}
)

// gettingSquareWG считает квадрат числа и выводит его
func gettingSquareWG(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintf(OUT, "%d ", n*n)
}

/*
gettingSquareChan считает квадрат числа и помещает его в канал squares,
после чего отчитывается о выполнении в канал isCompleted
*/
func gettingSquareChan(n int, squares chan<- int, isCompleted chan<- struct{}) {
	squares <- n * n
	isCompleted <- struct{}{}
}

func gettingSquareChanNotBuffChan(n int, squares chan<- int) {
	squares <- n * n
}

func main() {
	// через WaitGroup
	{
		wg := &sync.WaitGroup{}
		for _, v := range DATA {
			wg.Add(1)
			go gettingSquareWG(v, wg)
		}
		wg.Wait()
	}
	fmt.Fprintln(OUT)

	//через каналы
	{
		squares := make(chan int, len(DATA))

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
			fmt.Fprintf(OUT, "%d ", sq)
		}
	}
}
