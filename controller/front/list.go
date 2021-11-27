package front

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"saas/driver"
)


type List struct {
}

func (h List) List(c *gin.Context) {
	err := driver.RDB.Set(ctx,"change","hello world",0).Err()
	if err != nil {
		log.Fatal(err)
	}

	val,err := driver.RDB.Get(ctx,"change").Result()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"data":val,
	})
}
