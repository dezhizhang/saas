package back

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type FocusController struct {
	BaseController
}


func (f FocusController) Home(c *gin.Context) {
	c.HTML(http.StatusOK,"back/focus/index.html",nil)
}

func (f FocusController) Add(c *gin.Context) {
	c.HTML(http.StatusOK,"back/focus/add.html",nil)
}