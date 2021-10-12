package front

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Home struct {

}

func (h Home) Home(c *gin.Context)  {
	c.String(http.StatusOK,"hello")
}