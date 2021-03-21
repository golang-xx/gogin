package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	router := gin.Default()

	router.GET("/app",func(ctx *gin.Context){
		ctx.JSON(http.StatusOK,gin.H{
			"code" : 0,
			"message" : "success2",
		})
	})

	router.Run(":8080")
	return
}