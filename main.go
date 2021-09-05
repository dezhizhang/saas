package main

import (
	"github.com/gin-gonic/gin"
	"saas/router"
)

func main() {
	r := gin.Default()

	router.Back(r)

	r.Run()


}