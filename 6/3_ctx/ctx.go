package main

import (
	"context"
	"fmt"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

/*
Вариант 3 - с помощью контекста (по факту те же каналы)
*/

func worker(ctx context.Context, isDone chan<- struct{}) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker's job here is done")
			isDone <- struct{}{}
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	// по факту может использоваться любой контекст с поддержкой функции cancel
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	isDone := make(chan struct{})
	go worker(ctx, isDone)
	<-isDone
}
