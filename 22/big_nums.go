package main

import (
	"fmt"
	"math/big"
	"unsafe"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две
числовых переменных a,b, значение которых > 2^20.
*/

/*
Формально, весь функционал уже представлен в пакете math/big. Но можно написать свою оболочку над стандартными методами
*/
func multiply(a, b big.Int) *big.Int {
	var res big.Int
	return res.Mul(&a, &b)
}

func divide(a, b big.Int) *big.Int {
	var res big.Int

	// еще есть res.Quo
	return res.Div(&a, &b)
}

func add(a, b big.Int) *big.Int {
	var res big.Int
	return res.Add(&a, &b)
}

func subtract(a, b big.Int) *big.Int {
	var res big.Int
	return res.Sub(&a, &b)
}

func main() {
	var a, b big.Int
	a.SetBit(&a, 200, 1)
	b.SetString("99999999999999999999", 10)

	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b))
	fmt.Printf("%d\n", &a)
	fmt.Printf("%d\n", &b)

	fmt.Println()

	fmt.Printf("%d\n", multiply(a, b))
	fmt.Printf("%d\n", divide(a, b))
	fmt.Printf("%d\n", add(a, b))
	fmt.Printf("%d\n", subtract(a, b))
}
