package driver

import (
	"fmt"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)
var DB *gorm.DB
var err error

func init() {
	// 读取数据库配置
	config,err := ini.Load("./config/config.ini")
	if err != nil {
		log.Fatal(err)
		return
	}

	ip := config.Section("mysql").Key("ip").String()
	port := config.Section("mysql").Key("port").String()
	user := config.Section("mysql").Key("user").String()
	//password := config.Section("mysql").Key("password").String()
	database := config.Section("mysql").Key("database").String()
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",user,"sdf@df%%$65#fdsbXT",ip,port,database)
	//dsn := "root:sdf@df%%$65#fdsbXT@tcp(127.0.0.1:3306)/saas?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("连接成功")
}

