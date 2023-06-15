package gins

import "github.com/gin-gonic/gin"

var app *gin.Engine

func NewApp(ms ...func(ctx *gin.Context)) *gin.Engine {
	app = gin.New()
	if len(ms) == 0 {
		ms = append(ms, gin.Recovery())
	}
	for _, m := range ms {
		app.Use(m)
	}
	return app
}

func GetApp() *gin.Engine {
	return app
}
