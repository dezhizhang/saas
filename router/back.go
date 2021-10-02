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
		v1.GET("/role",back.RoleController{}.Home)
		v1.GET("/role/add",back.RoleController{}.Add)
		v1.POST("/role/doAdd",back.RoleController{}.DoAdd)
		v1.GET("/role/edit",back.RoleController{}.Edit)
		v1.POST("/role/doEdit",back.RoleController{}.DoEdit)
		v1.GET("/role/delete",back.RoleController{}.Delete)

	}
}