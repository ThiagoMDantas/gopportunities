package handler

import (
	"fmt"
	"net/http"

	"github.com/ThiagoMDantas/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json: "message"`
	ErrorCode string `json: "errorCode"`
}

type CreateOpeningResponse struct {
	Message string `json: "message"`
	Data    schemas.OpeningResponse
}

type DeleteOpeningResponse struct {
	Message string `json: "message"`
	Data    schemas.OpeningResponse
}

type ShowOpeningResponse struct {
	Message string `json: "message"`
	Data    schemas.OpeningResponse
}

type ListOpeningResponse struct {
	Message string `json: "message"`
	Data    []schemas.OpeningResponse
}

type UpdateOpeningResponse struct {
	Message string `json: "message"`
	Data    schemas.OpeningResponse
}
