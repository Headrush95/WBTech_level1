package main

import (
	"fmt"
)

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
собственное множество.
*/

// Set имплементация множества
type Set[T comparable] struct {
	val map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{val: make(map[T]struct{})}
}

func (s *Set[T]) Add(value T) {
	s.val[value] = struct{}{}
}
func (s *Set[T]) AddValues(values ...T) {
	for _, val := range values {
		s.Add(val)
	}
}

func (s *Set[T]) Print() {
	for val := range s.val {
		fmt.Print(val, " ")
	}
}

func main() {
	inputString := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println()
	set := NewSet[string]()
	set.AddValues(inputString...)
	set.Print()
}
