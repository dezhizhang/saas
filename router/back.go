package router

import (
	"github.com/gin-gonic/gin"
	"saas/controller/back"
)

func Back(r *gin.Engine) {
	v1 := r.Group("/admin")
	{
		v1.GET("/login",back.Login{}.Home)

		// 权限
		v1.GET("/role",back.RoleController{}.Home)
		v1.GET("/role/add",back.RoleController{}.Add)
		v1.POST("/role/doAdd",back.RoleController{}.DoAdd)
		v1.GET("/role/edit",back.RoleController{}.Edit)
		v1.POST("/role/doEdit",back.RoleController{}.DoEdit)
		v1.GET("/role/delete",back.RoleController{}.Delete)

		// 管理员
		v1.GET("/manager",back.ManagerController{}.Home)
		v1.GET("/manager/add",back.ManagerController{}.Add)
		v1.POST("/manager/doAdd",back.ManagerController{}.DoAdd)
		v1.GET("/manager/edit",back.ManagerController{}.Edit)
		v1.POST("/manager/doEdit",back.ManagerController{}.DoEdit)
		v1.GET("/manager/delete",back.ManagerController{}.Delete)

		// 轮播图
		v1.GET("/focus",back.FocusController{}.Home)
		v1.GET("/focus/add",back.FocusController{}.Add)
		v1.POST("/focus/doAdd",back.FocusController{}.DoAdd)

	}
}