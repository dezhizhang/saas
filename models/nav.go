package models


type Nav struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Url string `json:"url"`
	Status string `json:"status"`
	Sort int `json:"sort"`
}