package handlers

import (
	"github.com/gin-gonic/gin"
)

func PutOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"teste": "teste",
	})
}
