package main

import (
	"fmt"
	"math/rand"
	"strings"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как
это исправить? Приведите корректный пример реализации.

var justString string
func someFunc() {
v := createHugeString(1 << 10)
justString = v[:100]
}
func main() {
someFunc()
}
*/

/*
	Во-первых, создается излишне длинная строка в виде непонятного посредника v.
	Во-вторых, значение присваивается глобальной переменной, то есть строка будет занимать память все время работы
программы, пока ее значение не изменят (но, возможно, того требует конкретная задача).
	В третьих, имя глобальной переменной лучше писать отлично от локальных - все прописные буквы или с заглавной, например.
	В четвертых, чтобы избежать выделения лишней памяти, лучше передавать параметр размера требуемой строки в someFunc().
Тогда отпадает необходимость в создании подстроки (тоже могут быть проблемы, так как не сказано, что именно творится в
createHugeString() и при использовании символов, отличных от ASCII, могут возникнут артефакты).
	В пятых, переменную justString лучше сделать указателем на строку, а из createHugeString() возвращать указатель,
так как подразумевается возврат "объемной" строки и ее копирование будет дорогим по памяти.
	В шестых, не понятно предназначения someFunc, так как в ней просто меняется значение глобальной переменной пакета,
если сделать функцию видимой для остальных пакетов, то может быть в этом будет какой-то профит.
*/

var JustString *string

// ориентировочно что происходит в createHugeString()
func createHugeString(size int) *string {
	// поскольку используются символы, отличные от ASCII (кириллица), тип данных - rune
	refSet := []rune("qwertyuiopasdfghjklzxcvbnm1234567890-=<>,./?'\"\\[]{}!@#$%^&*()`~№;:йцукенгшщзхъфываапролджэячсмитьбю")
	var sb strings.Builder

	for i := 0; i < size; i++ {
		sb.WriteRune(refSet[rand.Intn(len(refSet))])
	}
	res := sb.String()
	return &res
}

func SomeFunc(size int) {
	JustString = createHugeString(size)
}
func main() {
	SomeFunc(100)
	fmt.Println(*JustString)
}