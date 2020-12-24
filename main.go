package main

import (
	"context"
	"github.com/google/logger"
	"go.uber.org/dig"
	"net/http"
	"os"
	"os/signal"
	"prime-data/api"
	"prime-data/app"
	"prime-data/pkg/jwt"
	"prime-data/router"
	"prime-data/services"
	"syscall"
	"time"
)

func main()  {
	container := dig.New()
	authen, err := app.InitAuth()
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

	engine := router.InitGinEngine(container)

	server := &http.Server{
		Addr:    ":8888",
		Handler: engine,
	}

	go func() {
		// service connections
		logger.Info("Listen at:", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Error: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown: ", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		logger.Info("Timeout of 5 seconds.")
	}
	logger.Info("Server exiting")
}

