package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpeningHandle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET opening",
	})
}
