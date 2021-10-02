package models


type Manager struct {
	Id int `json:"id"`
	RoleId int `json:"role_id"`
	Username string `json:"username"`
	Mobile string `json:"mobile"`
	Status int `json:"status"`
	Email string `json:"email"`
	AddTime int `json:"add_time"`
	IsSuper int `json:"is_super"`
	Password string `json:"password"`
	Role Role `gorm:"foreignKey:RoleId"`
}