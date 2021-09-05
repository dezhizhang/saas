package router

import (
	"github.com/gin-gonic/gin"
	"saas/controller/back"
)

func Back(r *gin.Engine) {
	v1 := r.Group("/api/v1/back")
	v1.GET("/user",back.User{}.List)
	v1.POST("/user/add",back.User{}.Add)
}