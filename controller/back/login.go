package back

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"saas/utils"
)

type Login struct {

}

func (l Login) Home(c *gin.Context) {
	id,b64s,err := utils.Captcha()
	if err != nil {
		log.Fatal(err)
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"status":200,
		"message":"ok",
		"data":map[string]interface{}{
			"url":b64s,
			"id":id,
		},
	})

	//c.HTML(http.StatusOK,"back/login/index.html",gin.H{})
}