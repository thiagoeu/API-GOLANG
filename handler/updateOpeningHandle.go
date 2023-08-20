package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thiagoeu/GOstd/schemas"
)

func UpdateOpeningHandle(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validade(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	// updateopening

	if request.Role != "" {
		opening.Role = request.Role

	}
	if request.Company != "" {
		opening.Company = request.Company

	}
	if request.Location != "" {
		opening.Location = request.Location

	}
	if request.Link != "" {
		opening.Link = request.Link

	}
	if request.Remote != nil {
		opening.Remote = *request.Remote

	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	// save opening

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error update opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}

	sendSucess(ctx, "update-opening", opening)
}
