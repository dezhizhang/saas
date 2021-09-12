package router

import (
	"github.com/gin-gonic/gin"
	"saas/controller/front"
)

func Front(r *gin.Engine) {

	r.GET("/home",front.Home{}.Home)

}