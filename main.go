package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"saas/router"
	"saas/utils"
)

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"UnixToTime":utils.UnixToTime,
	})
	r.LoadHTMLGlob("templates/**/**/*")
	r.Static("/static", "./static")
	router.Front(r)
	router.Back(r)

	r.Run()

}