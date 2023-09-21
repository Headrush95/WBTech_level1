package main

import (
	"cmp"
	"fmt"
	"math/rand"
	"slices"
)

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами
языка.
*/

// рандомизируем выбор среднего значения в слайсе
func selectDivider[T cmp.Ordered](arr []T) int {
	return rand.Intn(len(arr) - 1)
}

// quickSort осуществляет быструю сортировку слайса на месте
func quickSort[T cmp.Ordered](arr []T) {
	// если в рассматриваемом куске массива один элемент, то сортировка окончена
	if len(arr) < 2 {
		return
	}
	dividerIndex := selectDivider(arr) // индекс среднего значения, относительно которого будем сортировать на две половины исходный слайс
	divider := arr[dividerIndex]

	lower := 0
	higher := len(arr) - 1
loop:
	for {
		// ищем с конца слайса элемент, который будет меньше divider
		for arr[higher] > divider {
			higher--
			if higher <= lower {
				break loop
			}
		}
		// ищем с начала слайса элемент, который будет больше divider
		for arr[lower] < divider {
			lower++
			if lower >= higher {
				break loop
			}
		}

		// так как arr[lower] соответствует элемент, больший divider, а arr[higher] - меньший, меняем их местами
		arr[lower], arr[higher] = arr[higher], arr[lower]
		// проверяем, не переместился ли сам divider, если так, то обновляем его индекс
		if dividerIndex == lower {
			dividerIndex = higher
		} else if dividerIndex == higher {
			dividerIndex = lower
		}
	}

	// рекурсивно вызываем быструю сортировку двух половин: с элементами меньше, чем divider, и с элементами больше
	quickSort(arr[0:dividerIndex])
	quickSort(arr[dividerIndex+1:])
}

func main() {
	input := []int{3, 2, 10, 8, 1, 5, 9, 7, 6, 4}
	check := make([]int, len(input))
	copy(check, input)
	// наша сортировка
	quickSort(input)
	// сортировка из стандартного набора (есть еще sort.Sort(data))
	slices.Sort(check)
	fmt.Println(slices.Equal(check, input))
}
