package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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

	store := cookie.NewStore([]byte("secret111"))

	r.Use(sessions.Sessions("sessions",store))
	router.Front(r)
	router.Back(r)

	r.Run()

}