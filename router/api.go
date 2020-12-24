package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/logger"
	"go.uber.org/dig"
	"prime-data/api"
	"prime-data/middleware"
	"prime-data/pkg/http/wrapper"
	"prime-data/pkg/jwt"
)

func RegisterAPI(r *gin.Engine, container *dig.Container) error {
	err := container.Invoke(func(
		jwt jwt.IJWTAuth,
		auth *api.Auth,
	) error {
		{
			r.POST("/api/login", wrapper.Wrap(auth.Login))
		}

		api := r.Group("/app", middleware.UserAuthMiddleware(jwt))
		{
			api.GET("/hello", wrapper.Wrap(auth.Hello))
			api.GET("/request", )
		}
		return nil
	})

	if err != nil {
		logger.Error(err)
	}

	return err
}