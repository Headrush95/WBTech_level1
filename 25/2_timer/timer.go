package main

import "time"

/*
Реализовать собственную функцию sleep
*/

/*
Вариант 2 - с использованием таймера
*/

func sleep(d time.Duration) {
	// time.After под капотом использует time.Timer
	ticker := time.After(d)
	<-ticker
}

func main() {
	sleep(5 * time.Second)
}
