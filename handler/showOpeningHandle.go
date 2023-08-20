package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thiagoeu/GOstd/schemas"
)

func ShowOpeningHandle(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
	}

	sendSucess(ctx, "show opening", opening)

}
