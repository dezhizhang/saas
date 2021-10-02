
package models


type Focus struct {
	Id int `json:"id"`
	Img string `json:"img"`
	Type int `json:"type"`
	Name string `json:"Name"`
	Link string `json:"link"`
	Status string `json:"status"`
	Sort int `json:"sort"`
	AddTime int `json:"add_time"`
}