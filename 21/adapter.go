package main

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере
*/

// Adapter - реализует интерфейс клиента ClientGetUserInfo
type Adapter struct {
	service OurGetUserInfo
}

func NewAdapter(service OurGetUserInfo) *Adapter {
	return &Adapter{service: service}
}

func (a *Adapter) GetName() string {
	return fmt.Sprintf("%s %s", a.service.GetFirstName(), a.service.GetLastName())
}

func (a *Adapter) GetUsername() string {
	return a.service.GetNickName()
}

func (a *Adapter) GetEmail() string {
	return a.service.GetEmail()
}

func (a *Adapter) getPassword() string {
	return getPasswordFromHash(a.service.getPasswordHash())
}

func getPasswordFromHash(passwordHash string) string {
	return passwordHash
}
