package models

type ResponseUsers struct {
	Users []*User `json:"users"`
}

type User struct {
	Id             int    `json:"id,int"`
	PassportSeries int    `json:"passport_series,int"`
	PassportNumber int    `json:"passport_number,int"`
	Surname        string `json:"surname"`
	Name           string `json:"name"`
	Patronymic     string `json:"patronymic,omitempty"`
	Address        string `json:"address"`
}
