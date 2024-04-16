package api

import "github.com/gin-gonic/gin"

func SayA(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello a!",
	})
}

func SayB(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello b!",
	})
}
