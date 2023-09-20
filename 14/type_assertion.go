package main

import "fmt"

/*
Разработать программу, которая в рантайме способна определить тип
переменной: int, string, bool, channel из переменной типа interface{}.
*/

func checkType(value any) {
	switch value.(type) {
	case int:
		fmt.Println("Hey, it's integer")
	case string:
		fmt.Println("Hey, it's string")
	case bool:
		fmt.Println("Hey, it's boolean")
	case chan int:
		fmt.Println("Hey, it's channel")
	case chan string:
		fmt.Println("Hey, it's channel")
	default:
		fmt.Println("Who are you?")
	}
}

func main() {
	checkType(10)
	checkType("")
	checkType(false)
	checkType(make(chan int))
	type unknown struct{}
	checkType(unknown{})
}
