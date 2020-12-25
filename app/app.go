package app

import (
	"github.com/google/logger"
	"go.uber.org/dig"

	"prime-data/api"
	"prime-data/pkg/jwt"
	"prime-data/services"
)

func BuildContainer() *dig.Container {
	container := dig.New()
	authen, err := InitAuth()
	_ = container.Provide(func() jwt.IJWTAuth {
		return authen
	})

	err = services.Inject(container)
	if err != nil {
		logger.Error("Failed to inject services", err)
	}

	err = api.Inject(container)
	if err != nil {
		logger.Error("Failed to inject APIs", err)
	}

	return container
}