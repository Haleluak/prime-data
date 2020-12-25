package router

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func InitGinEngine(container *dig.Container, e *casbin.Enforcer) *gin.Engine {
	app := gin.New()
	Docs(app)
	RegisterAPI(app, container, e)
	return app
}
