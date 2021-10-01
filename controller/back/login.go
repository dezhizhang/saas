package back

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"saas/driver"
	"saas/models"
	"saas/utils"
)

type Login struct {

}

func (l Login) Home(c *gin.Context) {
	c.HTML(http.StatusOK,"back/login/index.html",gin.H{})
}

func (l Login) DoLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var manager []models.Manager
	driver.DB.Where("username=? AND password?=",username,	utils.Md5(password)).Find(&manager)

	if len(manager) > 0 {
		session := sessions.Default(c)
		session.Set("userInfo",&manager)
		session.Save()

	}


}