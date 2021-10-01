package models


type Manager struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Age int `json:"age"`
	Email string `json:"email"`
	AddTime int `json:"add_time"`
}