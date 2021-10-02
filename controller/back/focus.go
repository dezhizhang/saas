package back

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"saas/driver"
	"saas/models"
	"saas/utils"
)

type FocusController struct {
	BaseController
}


func (f FocusController) Home(c *gin.Context) {
	var focus []models.Focus
	driver.DB.Find(&focus)
	c.HTML(http.StatusOK,"back/focus/index.html",gin.H{
		"focus":focus,
	})
}

func (f FocusController) Add(c *gin.Context) {
	c.HTML(http.StatusOK,"back/focus/add.html",nil)
}

func (f FocusController) DoAdd(c *gin.Context) {
	pathName := "/admin/focus/add"
	name := c.PostForm("name")
	link := c.PostForm("link")
	focusType,err :=utils.Int(c.PostForm("focus_type"))
	if err != nil {
		f.Error(c,"图片类型错误",pathName)
		return
	}

	img,err := utils.UploadFile(c,"img")
	if err != nil {
		f.Error(c,"上传文件错误",pathName)
		return
	}

	focus :=  models.Focus{
		Name: name,
		Link: link,
		FocusType: focusType,
		Img: img,
		Status: 1,
		Sort: 1,
		AddTime: int(utils.GetUnix()),
	}

	err = driver.DB.Create(&focus).Error
	if err != nil {
		f.Error(c,"新增轮播图失败",pathName)
		return
	}
	f.Success(c,"新增轮播成功","/admin/focus")

}

func (f FocusController) Delete(c *gin.Context) {
	pathName := "/admin/focus"
	id,err :=utils.Int(c.Query("id"))
	if err != nil{
		f.Error(c,"参数错误",pathName)
		return
	}
	var focus models.Focus
	err = driver.DB.Where("id=?",id).Delete(&focus).Error
	if err != nil {
		f.Error(c,"删除轮播图失败",pathName)
		return
	}
	f.Success(c,"删除轮播图成功",pathName)
}