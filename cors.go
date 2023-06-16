package gins

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "POST,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "*")
		method := ctx.Request.Method
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}
	}
}
