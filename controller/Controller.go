package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(context *gin.Context){

	context.JSON(http.StatusOK, gin.H{
		"title": "GoAPIs",
		"message" : "Create API With Gin Golang",
	})
}
