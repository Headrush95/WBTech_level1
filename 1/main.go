package main

import (
	"errors"
	"fmt"
	"log"
)

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры
Human (аналог наследования).
*/

var (
	invalidAge  = errors.New("age cannot be negative")
	invalidName = errors.New("name cannot be empty")
)

type Human struct {
	name string
	age  int
	sex  string
}

func NewHuman(name string, age int, sex string) (*Human, error) {
	if age < 0 {
		return nil, invalidAge
	}
	if name == "" {
		return nil, invalidName
	}
	if sex != "male" && sex != "female" {
		return nil, errors.New("sex must be male or female")
	}
	return &Human{name: name, age: age, sex: sex}, nil
}

func (h *Human) Greet() {
	fmt.Println("Hello, my name is", h.name)
}

func (h *Human) SaySomething(text string) {
	fmt.Printf("[%s] %s\n", h.name, text)
}

func (h *Human) StealthyRoar() {
	fmt.Println("ROAR!!")
}

// встраиваем структуру Human в структуру Action
type Action struct {
	*Human
	status string
}

func NewAction(human *Human) *Action {
	return &Action{Human: human}
}

func (a *Action) Walk() {
	a.status = "walking"
	fmt.Printf("[%s] I'm walking\n", a.name)
}

func (a *Action) Eat() {
	a.status = "eating"
	fmt.Printf("[%s] I'm eating\n", a.name)
}

func main() {
	Ilya, err := NewHuman("Ilya", 27, "male")
	if err != nil {
		log.Fatalln(err)
	}

	Ilya.Greet()
	Ilya.SaySomething("Wow, new function!")
	IlyasActions := NewAction(Ilya)
	IlyasActions.Walk()
	IlyasActions.Eat()
	IlyasActions.Greet()
	IlyasActions.StealthyRoar()
}
