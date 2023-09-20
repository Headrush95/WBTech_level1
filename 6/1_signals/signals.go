package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

/*
Вариант 1 - с помощью сигнального канала
*/

func main() {
	quit := make(chan os.Signal, 1)
	defer close(quit)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	for {
		select {
		case <-quit:
			fmt.Println("Program is closing...")
			return
		default:
			fmt.Println("Doing something")
			time.Sleep(time.Second)
		}
	}
}
