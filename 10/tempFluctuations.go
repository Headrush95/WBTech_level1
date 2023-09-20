package main

import "fmt"

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
градусов. Последовательность в подмножноствах не важна.
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

func main() {
	tempGroups := make(map[int][]float64, 10)
	data := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	var group int
	for _, temp := range data {
		group = int(temp/10) * 10
		tempGroups[group] = append(tempGroups[group], temp)
	}

	// думаю, формат вывода данных тут не так важен
	for k, v := range tempGroups {
		fmt.Printf("%d: %v, ", k, v)
	}
}
