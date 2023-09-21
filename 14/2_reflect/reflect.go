package main

import (
	"fmt"
	"reflect"
)

/*
Разработать программу, которая в рантайме способна определить тип
переменной: int, string, bool, channel из переменной типа interface{}.
*/

/*
Вариант 2 - используя пакет reflect
*/

func main() {
	fmt.Println(reflect.TypeOf(10))
	fmt.Println(reflect.TypeOf("10"))
	fmt.Println(reflect.TypeOf(10.2))
	fmt.Println(reflect.TypeOf(make(chan int)))
	fmt.Println(reflect.TypeOf(struct{}{}))
	fmt.Println(reflect.TypeOf(false))

}
