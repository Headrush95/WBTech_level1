package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

/*
Вариант 2 - с помощью каналов
*/

// closer посылает через 4 секунды сигнал на завершение worker
func closer(done chan<- struct{}) {
	time.Sleep(4 * time.Second)
	done <- struct{}{}
}

// worker что-то делает, пока не поступит сигнал от closer
func worker(nums chan<- int, done <-chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("Worker is closing")
			// даем сигнал anotherWorker, что больше чисел не будет
			close(nums)
			return
		default:
			fmt.Println("Working hard...")
			nums <- rand.Intn(100)
			time.Sleep(time.Second)
		}
	}
}

// anotherWorker показывает способ завершения горутины с помощью проверки закрытия канала
func anotherWorker(nums <-chan int, quit chan<- struct{}) {
	for {
		num, ok := <-nums
		if !ok {
			fmt.Println("AnotherWorker is closing...")
			// даем главной горутине сигнал к завершению
			close(quit)
			return
		}
		fmt.Printf("Recieved number: %d\n", num)
	}
}

func main() {
	// done для завершения работы воркера
	done := make(chan struct{}, 1)
	defer close(done)
	go closer(done)
	// quit небуфферизированный для выхода из главной горутины
	quit := make(chan struct{})
	nums := make(chan int, 1)
	go worker(nums, done)
	go anotherWorker(nums, quit)

	//блокируем завершение main, пока не отработают остальные горутины
	<-quit
}
