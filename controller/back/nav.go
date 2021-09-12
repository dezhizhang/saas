package back

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"saas/driver"
	"saas/models"
)

type  Nav struct {

}

func (n Nav) List(c *gin.Context)  {
	var navList []models.Nav
	id := c.GetInt("id")
	fmt.Println("id",id)

	//driver.DB.Where("id > ?",3).Find(&navList)
	driver.DB.Where("id > ? and id < ?",1,4).Find(&navList)
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"ok",
		"success":true,
		"data":navList,
	})
}