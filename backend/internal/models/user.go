package models

type User struct {
	Id             int    `json:"id"`
	PassportSeries int    `json:"passport_series"`
	PassportNumber int    `json:"passport_number"`
	Surname        string `json:"surname"`
	Name           string `json:"name"`
	Patronymic     string `json:"patronymic"`
	Address        string `json:"address"`
}
