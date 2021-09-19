package back

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Manager struct {

}

func (m Manager) Home(c *gin.Context) {
	c.HTML(http.StatusOK,"back/manager/index.html",gin.H{})
}