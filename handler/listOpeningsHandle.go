package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOpeningsHandle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "LIST opening",
	})
}
