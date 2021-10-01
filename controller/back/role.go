package back

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Role struct {

}

func (r Role) Home(c *gin.Context)  {
	c.HTML(http.StatusOK,"back/role/index.html",nil)
}

func (r Role) Add(c *gin.Context)  {

}

func (r Role) Edit(c *gin.Context) {

}

func (r Role) Delete(c *gin.Context) {

}