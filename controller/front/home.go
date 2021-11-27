package front

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

type Home struct {
}

func (h Home) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "front/home/index.html", gin.H{})
}
