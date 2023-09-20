package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/
func main() {
	num1 := 13
	num2 := 100
	num1, num2 = num2, num1
	fmt.Printf("num1: %d, num2: %d\n", num1, num2)
}
