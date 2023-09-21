package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке
уникальные (true — если уникальные, false etc). Функция проверки должна быть
регистронезависимой.
Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

// hasUniqueSymbols проверяет строку на уникальность всех символов. Регистр не учитывается
func hasUniqueSymbols(src string) bool {
	src = strings.ToLower(src)
	runes := []rune(src)
	uniqueSymbols := make(map[rune]struct{}, len(runes))
	for _, rn := range runes {
		if _, ok := uniqueSymbols[rn]; ok {
			return false
		}
		uniqueSymbols[rn] = struct{}{}
	}
	return true
}

func main() {
	fmt.Println(hasUniqueSymbols("abcd"))
	fmt.Println(hasUniqueSymbols("abCdefAaf"))
	fmt.Println(hasUniqueSymbols("aabcd"))

}
