package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleResponse(ctx *gin.Context, code int, message string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message": message,
	})
}

func handleResponseWithData(ctx *gin.Context, code int, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"data": data,
	})
}

func OnBadRequest(ctx *gin.Context, message string) {
	handleResponse(ctx, http.StatusBadRequest, message)
}

func OnCreated(ctx *gin.Context, data interface{}) {
	handleResponseWithData(ctx, http.StatusCreated, data)
}

func OnError(ctx *gin.Context, message string) {
	handleResponse(ctx, http.StatusInternalServerError, message)
}

func OnNotFound(ctx *gin.Context, message string) {
	handleResponse(ctx, http.StatusNotFound, message)
}

func OnSuccessWithData(ctx *gin.Context, data interface{}) {
	handleResponseWithData(ctx, http.StatusOK, data)
}

func OnSuccess(ctx *gin.Context, message string) {
	handleResponse(ctx, http.StatusOK, message)
}
