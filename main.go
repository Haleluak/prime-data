package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"prime-data/models"
	"syscall"
	"time"

	"github.com/google/logger"
	"prime-data/app"
	"prime-data/router"
)

func main() {
	container := app.BuildContainer()

	models.InitOrmDB()
	defer models.Client.Close()

	app.Migrate(models.Client)
	authEnforcer := app.InitCasbin()

	engine := router.InitGinEngine(container, authEnforcer)

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

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
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
