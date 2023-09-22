package main

/*
Реализовать паттерн «адаптер» на любом примере
*/

// ClientGetUserInfo - интерфейс клиента
type ClientGetUserInfo interface {
	GetName() string
	GetUsername() string
	GetEmail() string
	getPassword() string
}

// ClientUser - структура пользователя в сервисе клиента. Как видно, поля этой структуры отличаются от наешго OurUser
type ClientUser struct {
	name     string
	username string
	email    string
	password string
}

func (u *ClientUser) GetName() string {
	return u.name
}

func (u *ClientUser) GetUsername() string {
	return u.username
}

func (u *ClientUser) GetEmail() string {
	return u.email
}

func (u *ClientUser) getPassword() string {
	return u.password
}
