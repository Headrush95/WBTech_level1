package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// getWorkersCount возвращает количество воркеров
func getWorkersCount() int {
	var workersCount int
	_, err := fmt.Scanln(&workersCount)
	if err != nil {
		log.Fatalln(err)
	}
	return workersCount
}

// getWorkersTokens возвращает заполненный канал токенов
func getWorkersTokens() chan struct{} {
	count := getWorkersCount()
	ch := make(chan struct{}, count)

	// заполняем канал токенами
	for i := 0; i < count; i++ {
		ch <- struct{}{}
	}
	return ch
}

func worker(token chan<- struct{}) {
	time.Sleep(time.Second)
	fmt.Println("Worker's job here is done")
	token <- struct{}{}
}

func main() {
	// токены воркеров: когда воркер начинает работать - берет токен из канала, когда завершает - возвращает
	workerTokens := getWorkersTokens()
	defer close(workerTokens)

	// завершение работы воркеров через сигналы завершения
	quit := make(chan os.Signal, 1)
	defer close(quit)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	/*
		Использованные UNIX сигналы (почему-то работают и в Windows...)
		SIGTERM - Сигнал завершения (сигнал по умолчанию для утилиты kill)
		SIGINT - Сигнал прерывания (Ctrl-C) с терминала (тоже самое, что и os.Interrupt)
		SIGQUIT - Сигнал «Quit» с терминала (Ctrl-\)
	*/

	for {
		select {
		case <-workerTokens:
			go worker(workerTokens)
		case <-quit:
			fmt.Println("Exiting...")
			return // можно еще через break <loop label> (если после цикла есть какая-то логика)
		}
	}

}
