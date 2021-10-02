package back

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"saas/driver"
	"saas/models"
	"saas/utils"
	"strings"
)

type RoleController struct {
  BaseController
}

func (r RoleController) Home(c *gin.Context)  {
	var roles []models.Role
	driver.DB.Find(&roles)
	c.HTML(http.StatusOK,"back/role/index.html",gin.H{
		"roles":roles,
	})
}

func (r RoleController) Add(c *gin.Context)  {
	c.HTML(http.StatusOK,"back/role/add.html",nil)
}

func (r RoleController) DoAdd(c *gin.Context) {
	title := strings.Trim(c.PostForm("title"),"")
	description := strings.Trim(c.PostForm("description"),"")
	if title == "" {
		r.Error(c,"标题不能为空","/admin/role/add")
	}
	var role models.Role
	role.Title = title
	role.Description = description
	role.Status = 1
	role.AddTime = int(utils.GetUnix())

	err := driver.DB.Create(&role).Error
	if err != nil {
		r.Error(c,"增加角色失败,请重试","/admin/role/add")
		return
	}
	r.Success(c,"增加角色成功","/admin/role")

}

func (r RoleController) Edit(c *gin.Context) {
	id := c.Query("id")
	n,err := utils.Int(id)
	if err != nil {
		r.Error(c,"参数错误","/admin/role")
		return
	}
	var role models.Role
	driver.DB.Where("id=?",n).Find(&role)
	c.HTML(http.StatusOK,"back/role/edit.html",gin.H{
		"role": role,
	})
}

func (r RoleController) DoEdit(c *gin.Context) {
	strId := strings.Trim(c.PostForm("id"),"")
	id,err := utils.Int(strId)
	if err != nil {
		r.Error(c,"参数错误","/admin/role/edit")
	}
	title := strings.Trim(c.PostForm("title"),"")
	description := strings.Trim(c.PostForm("description"),"")

	if title == "" {
		r.Error(c,"标题不能为空","/admin/role/edit")
		return
	}

	var role models.Role
	driver.DB.Where("id=?",id).Find(&role)

	// 修改数据
	role.Title = title
	role.Description = description

	// 保存数据
	err = driver.DB.Save(&role).Error
	if err != nil {
		r.Error(c,"修改数据失败","/admin/role/edit?id="+utils.String(id))
		return
	}
	r.Success(c,"修改数据成功","/admin/role")

}

func (r RoleController) Delete(c *gin.Context) {

}