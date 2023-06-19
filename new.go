package gins

import (
	"github.com/gin-gonic/gin"
)

type Context = gin.Context
type RouterGroup = gin.RouterGroup

var app *Application

type Application struct {
	*gin.Engine
	urlpatterns map[string]gin.HandlerFunc
}

func New(ms ...func(ctx *Context)) *Application {
	app = &Application{
		Engine:      gin.New(),
		urlpatterns: map[string]gin.HandlerFunc{},
	}

	app.Use(gin.Recovery())
	for _, m := range ms {
		app.Use(m)
	}
	return app
}

func Register(url string, h gin.HandlerFunc) *Application {
	app.urlpatterns[url] = h
	return app
}
