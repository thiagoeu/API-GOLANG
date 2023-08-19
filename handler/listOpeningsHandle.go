package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thiagoeu/GOstd/schemas"
)

func ListOpeningsHandle(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSucess(ctx, "list-openings", openings)
}
