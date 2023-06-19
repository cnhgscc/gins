package gins

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Require(methods ...string) func(ctx *gin.Context) {
	_require := map[string]bool{}
	for _, s := range methods {
		_require[s] = true
	}
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		if _, ok := _require[method]; !ok {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}
	}
}
