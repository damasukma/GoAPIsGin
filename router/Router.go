package router

import (
	"github.com/gin-gonic/gin"
	"GoAPIs/controller"
	)



func SetupRouter() *gin.Engine{
	router := gin.Default()
	router.GET("/", controller.Index)

	router.Run()
	return router
}