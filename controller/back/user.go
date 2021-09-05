package back

import (
	"github.com/gin-gonic/gin"
	"log"
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

func (u User) Add(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Fatal(err)
	}
	driver.DB.Create(&user)
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"增加数据成功",
		"success":true,
		"data":user,
	})
}