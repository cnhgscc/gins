package gins

import "github.com/gin-gonic/gin"

func NewApp(ms ...func(ctx *gin.Context)) {
	app := gin.New()
	if len(ms) == 0 {
		ms = append(ms, gin.Recovery())
	}
	for _, m := range ms {
		app.Use(m)
	}
}
