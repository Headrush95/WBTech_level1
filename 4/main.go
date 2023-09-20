package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать
набор из N воркеров, которые читают произвольные данные из канала и
выводят в stdout. Необходима возможность выбора количества воркеров при
старте.
*/

// getWorkersCount возвращает количество воркеров
func getWorkersCount() int {
	var workersCount int
	fmt.Print("Enter workers count: ")
	_, err := fmt.Scanln(&workersCount)
	if err != nil {
		log.Fatalln(err)
	}
	return workersCount
}

// getWorkersTokens возвращает заполненный канал токенов
func getWorkersTokens(count int) chan struct{} {
	ch := make(chan struct{}, count)
	// заполняем канал токенами
	for i := 0; i < count; i++ {
		ch <- struct{}{}
	}
	return ch
}

func worker(token chan<- struct{}, in <-chan int) {
	// что-то делаем...
	time.Sleep(time.Second)
	res, ok := <-in
	if !ok {
		// возвращаем токен, тем самым показываем, что воркер завершил свою работу
		token <- struct{}{}
		return
	}
	fmt.Printf("Worker's job here is done. Get %d\n", res)
	token <- struct{}{}
}

func main() {
	// получаем от пользователя кол-во воркеров
	count := getWorkersCount()
	// токены воркеров: когда воркер начинает работать - берет токен из канала, когда завершает - возвращает
	workerTokens := getWorkersTokens(count)
	defer close(workerTokens)

	// завершение работы воркеров через сигналы завершения
	quit := make(chan os.Signal, 1)
	defer close(quit)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	/*
		Использованные UNIX сигналы (теперь работают и в Windows...)
		SIGTERM - Сигнал завершения (сигнал по умолчанию для утилиты kill)
		SIGINT - Сигнал прерывания (Ctrl-C) с терминала (тоже самое, что и os.Interrupt)
	*/
	data := make(chan int, 1)

	// откидываем писателя в отдельную горутину
	go func() {
		defer close(data)

		writerQuit := make(chan os.Signal, 1)
		defer close(writerQuit)
		signal.Notify(writerQuit, syscall.SIGTERM, syscall.SIGINT)

		for {
			select {
			case <-writerQuit:
				fmt.Println("Writer's job here is done")
				return
			default:
				data <- rand.Intn(1000)
			}

		}
	}()

	for {
		select {
		case <-workerTokens:
			go worker(workerTokens, data)
		case <-quit:
			fmt.Println("Waiting for workers to complete...")
			// ждем, пока все активные воркеры доделают свою работу
			for i := 0; i < count; i++ {
				<-workerTokens
			}
			fmt.Println("Exiting...")
			return // можно еще через break <loop label> (если после цикла есть какая-то логика)
		}
	}

}
