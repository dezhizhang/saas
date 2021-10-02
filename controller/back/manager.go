package back

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"saas/driver"
	"saas/models"
	"saas/utils"
	"strings"
)

type ManagerController struct {
	BaseController
}

func (m ManagerController) Home(c *gin.Context) {
	var managers []models.Manager
	driver.DB.Preload("Role").Find(&managers)
	c.HTML(http.StatusOK,"back/manager/index.html",gin.H{
		"managers":managers,
	})
}

func (m ManagerController) Add(c *gin.Context)  {
	var roles []models.Role
	driver.DB.Find(&roles)
	c.HTML(http.StatusOK,"back/manager/add.html",gin.H{
		"roles":roles,
	})
}

func (m ManagerController) DoAdd(c *gin.Context) {
	pathName := "/admin/manager/add"
	roleId,err :=utils.Int(c.PostForm("role_id"))
	if err != nil {
		m.Error(c,"传入参数有误",pathName)
		return
	}
	email := strings.Trim(c.PostForm("email"),"")
	mobile:= strings.Trim(c.PostForm("mobile"),"")
	username := strings.Trim(c.PostForm("username"),"")
	password := strings.Trim(c.PostForm("password"),"")

	manager := models.Manager{
		Username: username,
		Password: utils.Md5(password),
		Email: email,
		Mobile: mobile,
		RoleId: roleId,
		Status: 1,
		AddTime: int(utils.GetUnix()),
	}

	err = driver.DB.Create(&manager).Error
	if err != nil {
		m.Error(c,"增加管理员失败",pathName)
		return
	}
	m.Success(c,"增加管理员成功","/admin/manager")
}


func (m ManagerController) Edit(c *gin.Context) {
	pathName := "/admin/manager/edit"
	id ,err := utils.Int(c.Query("id"))
	if err != nil{
		m.Error(c,"传入参数有误",pathName)
		return
	}
	var roles []models.Role
	var manager models.Manager
	driver.DB.Find(&roles)
	driver.DB.Where("id=?",id).Find(&manager)
	c.HTML(http.StatusOK,"back/manager/edit.html",gin.H{
		"roles":roles,
		"manager":manager,
	})
}

func (m ManagerController) DoEdit(c *gin.Context) {
	fmt.Println(c.PostForm("id"))
	id,err := utils.Int(c.PostForm("id"))
	pathName := "/admin/manager/edit?id="+ c.PostForm("id")
	if err != nil {
		m.Error(c,"参数错误",pathName)
		return
	}

	roleId,_ := utils.Int(c.Param("role_id"))
	mobile := c.PostForm("mobile")
	email := c.PostForm("email")
	username := c.PostForm("username")
	password := c.PostForm("password")

	manager := models.Manager{
		Id: id,
		RoleId: roleId,
		Mobile: mobile,
		Email: email,
		Username: username,
		Password: password,
	}

	err = driver.DB.Save(&manager).Error
	if err != nil {
		m.Error(c,"修改管理失败",pathName)
		return
	}
	m.Success(c,"修改管理成功","/admin/manager")
}

func (m ManagerController) Delete(c *gin.Context) {
	pathName := "/admin/manager"
	id,err := utils.Int(c.Query("id"))
	if err != nil{
		m.Error(c,"参数错误",pathName)
		return
	}
	var manager models.Manager
	err = driver.DB.Where("id=?",id).Delete(&manager).Error
	if err != nil{
		m.Error(c,"删除失败",pathName)
		return
	}
	m.Success(c,"删除成功",pathName)
}