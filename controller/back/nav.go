package back

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"saas/models"
)

type  Nav struct {

}

func (n Nav) List(c *gin.Context)  {
	var navList []models.Nav
	id := c.GetInt("id")
	fmt.Println("id",id)

	//driver.DB.Where("id > ?",3).Find(&navList)
	//driver.DB.Where("id > ? and id < ?",1,4).Find(&navList)
	//driver.DB.Where("id in (?)",[]int{2,4}).Find(&navList)
	//driver.DB.Where("title like ? ","%科技%").Find(&navList)
	//driver.DB.Where("id between ? and ?",2,4).Find(&navList)
	//driver.DB.Where("id=? or id=?",2,3).Find(&navList)
	//driver.DB.Select("id,title").Find(&navList)
	//driver.DB.Order("id desc").Find(&navList)
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"ok",
		"success":true,
		"data":navList,
	})
}