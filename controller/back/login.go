package back

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {

}

func (l Login) Home(c *gin.Context) {
	c.HTML(http.StatusOK,"back/login/index.html",gin.H{})
}