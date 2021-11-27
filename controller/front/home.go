package front

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"saas/driver"
)

var ctx = context.Background()

type Home struct {

}

func (h Home) Home(c *gin.Context)  {
	err := driver.RDB.Set(ctx,"name","张三",1).Err()
	if err != nil {
		log.Fatal(err)
	}

	val,err := driver.RDB.Get(ctx,"name").Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("val",val)

	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"hello world",
	})
}