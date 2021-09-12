package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"saas/router"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	router.Front(r)
	router.Back(r)


	config,err := ini.Load("./config/config.ini")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println(config.Section("mysql").Key("password").String())

	r.Run()


}