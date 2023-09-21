package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

/*
Вариант через хэшмапу со счетчиком вхождений
*/

func intersection[T comparable](sets ...[]T) []T {
	/*
		поскольку в условиях говорится, что входные данные - множества,
		то проверку на дубликаты можно не делать
	*/
	intersectionCounter := make(map[T]int, len(sets[0])) // карта-счетчик для значений в множествах
	result := make([]T, 0, len(sets[0]))                 // результат пересечения
	for _, set := range sets {
		for _, value := range set {
			intersectionCounter[value]++
			if intersectionCounter[value] == len(sets) { // если счетчик равен кол-ву множеств, то значение есть во всех
				result = append(result, value)
			}
		}
	}
	return result
}
func main() {
	set1 := []int{0, 6, 3, 1, 5, 4}
	set2 := []int{8, 1, 6, 3, 4, 9, 10}
	fmt.Println(intersection(set1, set2))
	set3 := []string{"c", "b", "a", "d"}
	set4 := []string{"e", "d", "c", "f"}
	fmt.Println(intersection(set3, set4))
}
