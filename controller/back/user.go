package back

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"saas/driver"
	"saas/models"
)

type User struct {

}

func (u User) List(c *gin.Context) {
	var users []models.User
	driver.DB.Find(&users)
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"ok",
		"success":true,
		"data":users,
	})
}