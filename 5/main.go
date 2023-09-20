package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в
канал, а с другой стороны канала — читать. По истечению N секунд программа
должна завершаться.
*/

// getTimeOut устнавливает продолжительность работы программы
func getTimeOut() int {
	var timeout int
	_, err := fmt.Scanln(&timeout)
	if err != nil {
		log.Fatalln(err)
	}
	return timeout
}

func reader(in <-chan int) {
	fmt.Printf("Reading massege: %d\n", <-in)
	time.Sleep(500 * time.Millisecond)
}

func writer(out chan<- int, quit <-chan struct{}) {
	defer close(out)
	for {
		select {
		case <-quit:
			fmt.Println("writer is closing...")
			return
		default:
			out <- rand.Intn(100)
		}
	}
}

func main() {
	timeOut := getTimeOut()
	ticker := time.After(time.Duration(timeOut) * time.Second)
	dataChan := make(chan int, 1)

	// writerQuit сигнализирует писателю, что пора завершаться
	writerQuit := make(chan struct{}, 1)
	defer close(writerQuit)

	go writer(dataChan, writerQuit)
	for {
		select {
		case <-ticker:
			fmt.Println("Exiting...")
			writerQuit <- struct{}{}
			return
		default:
			reader(dataChan)
		}
	}

}
