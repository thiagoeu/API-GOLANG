package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thiagoeu/GOstd/schemas"
)

func CreateOpeningHandle(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validade(); err != nil {
		logger.Errorf("validade error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	oppening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&oppening).Error; err != nil {
		logger.Errorf("error creating opening %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}
	sendSucess(ctx, "Create Oppening", oppening)

}
