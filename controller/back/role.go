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
	c.HTML(http.StatusOK,"back/role/index.html",nil)
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
	r.Success(c,"增加角色成功","admin/role")

}

func (r RoleController) Edit(c *gin.Context) {

}

func (r RoleController) Delete(c *gin.Context) {

}