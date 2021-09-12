package back

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"saas/driver"
	"saas/models"
)

type  Nav struct {

}

func (n Nav) List(c *gin.Context)  {
	var navList []models.Nav
	driver.DB.Find(&navList)
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"ok",
		"success":true,
		"data":navList,
	})
}