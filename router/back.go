package router

import (
	"github.com/gin-gonic/gin"
	"saas/controller/back"
)

func Back(r *gin.Engine) {
	v1 := r.Group("/admin")
	{
		v1.GET("/login",back.Login{}.Home)
		v1.GET("/manager",back.Manager{}.Home)
	}
}