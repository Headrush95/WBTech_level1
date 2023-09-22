package main

/*
Реализовать паттерн «адаптер» на любом примере
*/

// OurGetUserInfo интерфейс нашего приложения
type OurGetUserInfo interface {
	GetFirstName() string
	GetLastName() string
	GetNickName() string
	GetEmail() string
	GetAddress() string
	getPasswordHash() string
}

// OurUser - структура пользователя в нашем сервисе, реализует интерфейс OurGetUserInfo
type OurUser struct {
	firstName    string
	lastName     string
	nickName     string
	email        string
	passwordHash string
	address      string
}

func (u *OurUser) GetFirstName() string {
	return u.firstName
}

func (u *OurUser) GetLastName() string {
	return u.lastName
}

func (u *OurUser) GetNickName() string {
	return u.nickName
}

func (u *OurUser) GetEmail() string {
	return u.email
}

func (u *OurUser) GetAddress() string {
	return u.address
}

func (u *OurUser) getPasswordHash() string {
	return u.passwordHash
}
