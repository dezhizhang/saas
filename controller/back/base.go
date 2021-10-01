package back

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {

}

func (b BaseController) success(c *gin.Context,message string,redirectUrl string) {
	c.HTML(http.StatusOK,"back/public/success.html",gin.H{
		"message":message,
		"redirectUrl":redirectUrl,
	})
}

func (b BaseController) error(c *gin.Context,message string,redirectUrl string)  {
	c.HTML(http.StatusOK,"back/public/error.html",gin.H{
		"message":message,
		"redirectUrl":redirectUrl,
	})
}
