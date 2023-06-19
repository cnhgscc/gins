package gins

import "github.com/gin-gonic/gin"

var app *Application

type Application struct {
	*gin.Engine
	urlpatterns map[string]gin.HandlerFunc
}

func New(ms ...func(ctx *gin.Context)) *Application {
	app = &Application{
		Engine:      gin.New(),
		urlpatterns: map[string]gin.HandlerFunc{},
	}

	if len(ms) == 0 {
		ms = append(ms, gin.Recovery())
	}
	for _, m := range ms {
		app.Use(m)
	}
	return app
}

func (app *Application) Register(url string, h gin.HandlerFunc) *Application {
	app.urlpatterns[url] = h
	return app
}
