package models


type Role struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Status int `json:"status"`
	Description string `json:"description"`
	AddTime int `json:"add_time"`
}