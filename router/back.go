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

		// 权限
		v1.GET("/role",back.Role{}.Home)
		v1.GET("/role/add",back.Role{}.Add)
		v1.GET("/role/edit",back.Role{}.Edit)
		v1.GET("/role/delete",back.Role{}.Delete)

	}
}