package router

import "github.com/gin-gonic/gin"

func SetupRouter() {
	r := gin.Default()
	r.POST("/info", Handle)
	r.Run()
}
