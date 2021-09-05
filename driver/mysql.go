package driver

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)
var DB *gorm.DB
var err error

func init() {
	dsn := "root:sdf@df%%$65#fdsbXT@tcp(127.0.0.1:3306)/saas?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("连接成功")
}

