package back

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"saas/driver"
	"saas/models"
	"saas/utils"
	"strings"
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


func (f FocusController) Edit(c *gin.Context) {
	id,err := utils.Int(c.Query("id"))
	if err != nil {
		f.Error(c,"参数错误","/admin/focus")
		return
	}
	var focus models.Focus
	driver.DB.Where("id=?",id).Find(&focus)
	c.HTML(http.StatusOK,"back/focus/edit.html",gin.H{
		"focus":focus,
	})
}

func (f FocusController) DoEdit(c *gin.Context)  {

	id,err := utils.Int(c.PostForm("id"))
	pathName := "/admin/focus/edit?id"+ c.PostForm("id")
	if err != nil {
		f.Error(c,"传入的参数有误",pathName)
		return
	}

	var focus models.Focus

	name := strings.Trim(c.PostForm("name"),"")
	link :=strings.Trim(c.PostForm("link"),"")
	focusType,err :=utils.Int(c.PostForm("focus_type"))
	if err != nil {
		f.Error(c,"类型错误",pathName)
		return
	}
	focus.Id = id
	focus.Name = name
	focus.Link = link
	focus.FocusType = focusType
	img,_ := utils.UploadFile(c,"img")
	if img != "" {
		focus.Img = img
	}

	err = driver.DB.Where("id=?",id).Save(&focus).Error
	if err != nil {
		f.Error(c,"修改数据失败",pathName)
		return
	}
	f.Success(c,"修改数据成功",pathName)

}