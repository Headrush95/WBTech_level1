package main

import (
	"fmt"
	"os"
)

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
массива, во второй — результат операции x*2, после чего данные из второго
канала должны выводиться в stdout.
*/

var (
	OUT  = os.Stdout
	DATA = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
)

// mult получает числа из канала in, умножает их на 2 и возвращает в канал out
func mult(in <-chan int, out chan<- int) {
	defer close(out)
	for num := range in {
		fmt.Fprintf(OUT, "[Mult] Get %d, sent %d\n", num, num*2)
		out <- num * 2
	}
}

// reader читает данные из переданного массива и отправляет их в канал
func reader(nums []int, out chan<- int) {
	defer close(out)
	for num := range nums {
		fmt.Fprintf(OUT, "[Reader] got %d\n", num)
		out <- num
	}
}

// writer пишет измененные числа в os.Stdout
func writer(in <-chan int, done chan struct{}) {
	defer close(done)
	for num := range in {
		fmt.Fprintf(OUT, "[Writer] Recieved number %d\n", num)
	}
}

func main() {
	// в justNums пишутся числа из массива DATA
	justNums := make(chan int, len(DATA)/2)
	// в mutatedNums пигутся числа после преобразования (x*2)
	mutatedNums := make(chan int, len(DATA)/2)

	// канал завершения, чтобы главная горутина не завершилась раньше времени.
	// Вместо канала можно использовать sync.WaitGroup
	done := make(chan struct{})

	go reader(DATA, justNums)
	go mult(justNums, mutatedNums)
	go writer(mutatedNums, done)

	// ждем, пока все данные запишуться в os.Stdout
	<-done
}
