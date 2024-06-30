package models

type User struct {
	Id         int    `json:"id"`
	Passport   string `json:"passport" example:"1234 567890"`
	Surname    string `json:"surname" example:"Ivanov"`
	Name       string `json:"name" example:"Ivan"`
	Patronymic string `json:"patronymic,omitempty" example:"Ivanovich"`
	Address    string `json:"address" example:"Moscow, Lenina 5 apt. 5"`
}

type ResponseUsers struct {
	Users []*User `json:"users"`
}

type UserPassport struct {
	PassportNumber string `json:"passport_number" example:"1234 567890"`
}
