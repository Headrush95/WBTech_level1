package main

import (
	"context"
	"time"
)

/*
Реализовать собственную функцию sleep
*/

/*
Вариант 1 - с использованием контекста
*/
func sleep(d time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), d)
	defer cancel()
	<-ctx.Done()
}

func main() {
	sleep(5 * time.Second)
}
