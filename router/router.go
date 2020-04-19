package router

import "github.com/gin-gonic/gin"

func SetupRouter() {
	r := gin.Default()
	r.GET("/info", Handle)
	r.Run()
}
