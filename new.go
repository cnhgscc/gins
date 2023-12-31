package gins

import (
	"github.com/gin-gonic/gin"
)

var app *Application

type Handler struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

func New(ms ...func(ctx *Context)) *Application {
	app = &Application{
		Engine:      gin.New(),
		urlpatterns: map[string]Handler{},
	}

	app.Use(gin.Recovery())
	for _, m := range ms {
		app.Use(m)
	}
	return app
}

func Register(method string, url string, h gin.HandlerFunc) *Application {
	app.urlpatterns[url] = Handler{
		Method:  method,
		Path:    url,
		Handler: h,
	}
	return app
}
