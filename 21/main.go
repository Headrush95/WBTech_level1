package main

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере
*/

func main() {
	adaptee := OurUser{
		firstName:    "Ilya",
		lastName:     "Makolin",
		nickName:     "headrush",
		email:        "simple@example.com",
		passwordHash: "kndfswe1283u4i1jrkwef",
		address:      "simple address",
	}

	adapter := NewAdapter(&adaptee)

	// теперь сторонний сервис может взаимодействовать через привычный ему интерфейс с нашей структурой пользователя
	fmt.Println(adapter.GetName())
}
