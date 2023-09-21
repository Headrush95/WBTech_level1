package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками, которые
представлены в виде структуры Point с инкапсулированными параметрами x,y и
конструктором
*/

type Point struct {
	x, y float64
}

func newPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func calcDistance(a, b *Point) float64 {
	// найдем относительные координаты
	zeroX := a.x - b.x
	zeroY := a.y - b.y
	//результат - гипотенуза между двумя точками
	return math.Sqrt(math.Pow(zeroX, 2) + math.Pow(zeroY, 2))
}

func main() {
	pointA := newPoint(2, 5)
	pointB := newPoint(7, 3)
	fmt.Println(calcDistance(pointA, pointB))
}
