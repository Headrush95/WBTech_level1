package main

import (
	"errors"
	"fmt"
	"log"
)

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в
1 или 0.
*/
// errorHandle вспомогательная функция, чтобы не плодить if err != nil...
func errorHandle(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	var num int64   // число для изменения
	var bit uint8   // изменяемый бит
	var value uint8 // новое значение изменяемого бита

	fmt.Print("Enter number: ")
	_, err := fmt.Scanln(&num)
	errorHandle(err)

	fmt.Print("Enter changing bit (starting from 0): ")
	_, err = fmt.Scanln(&bit)
	errorHandle(err)
	if bit > 63 {
		errorHandle(errors.New("invalid bit"))
	}

	fmt.Printf("Enter new %d bit value (0 or 1): ", bit)
	_, err = fmt.Scanln(&value)
	errorHandle(err)

	fmt.Printf("Old number %d (%b), ", num, num)
	if value == 1 {
		// устанавливаем 1 в указанный бит
		num |= 1 << bit
	} else if value == 0 {
		// устанавливаем 0 в указанный бит
		num &^= 1 << bit
	} else {
		errorHandle(errors.New("invalid bit value input"))
	}
	fmt.Printf("new number %d (%b)\n", num, num)
}
