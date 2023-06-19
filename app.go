package gins

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Context = gin.Context
type RouterGroup = gin.RouterGroup

type Application struct {
	*gin.Engine
	urlpatterns map[string]Handler
}

func (app *Application) Serve(host string, port int) {
	for _, h := range app.urlpatterns {
		switch h.Method {
		case http.MethodPost:
			app.POST(h.Path, h.Handler)
		case http.MethodGet:
			app.GET(h.Path, h.Handler)
		}
	}
	Serve(app, host, port)
}
