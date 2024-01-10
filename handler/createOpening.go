package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := Opening{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err!=nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Create(&request).Error; err != nil {
		logger.Errorf("Error creating opening: %v", err.Error())
	}

}