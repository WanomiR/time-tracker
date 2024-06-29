package models

type User struct {
	Id             int    `json:"id,int" example:"1"`
	PassportSeries int    `json:"passport_series,int" example:"1234"`
	PassportNumber int    `json:"passport_number,int" example:"567890"`
	Surname        string `json:"surname" example:"Ivanov"`
	Name           string `json:"name" example:"Ivan"`
	Patronymic     string `json:"patronymic,omitempty" example:"Ivanovich"`
	Address        string `json:"address" example:"Moscow, Lenina 5 apt. 5"`
}

type ResponseUsers struct {
	Users []*User `json:"users"`
}

type UserPassport struct {
	Series int `json:"series,int" example:"1234"`
	Number int `json:"number,int" example:"567890"`
}
