package middleware

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"saas/models"
	"strings"
)

func Auth(c *gin.Context) {
	// 获取url地址
	pathName := strings.Split(c.Request.URL.String(),"?")[0]

	// 获取session里面的信息
	session := sessions.Default(c)
	manager := session.Get("manager")
	userInfo,ok := manager.(string)
	if ok {
		var managerArr []models.Manager
		json.Unmarshal([]byte(userInfo),&managerArr)
		if !(len(managerArr) > 0 && managerArr[0].Username != "") {
			if pathName != "/admin/login" && pathName != "/admin/doLogin" {
				c.Redirect(302,"/admin/login")
			}
		}
	}else {
		if pathName != "/admin/login" && pathName != "/admin/doLogin" {
			c.Redirect(302,"/admin/login")
		}

	}

}