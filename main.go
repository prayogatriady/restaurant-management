package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/restaurant-management/api/lost"
	"github.com/prayogatriady/restaurant-management/http/router"
	"github.com/prayogatriady/restaurant-management/utils/config"
	"github.com/prayogatriady/restaurant-management/utils/db"
)

func init() {

	config.InitEnv()
	// logger.InitLogger()
	db.InitMySql()

}

func main() {

	appConfig := config.AppCfg

	app := gin.Default()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	app.NoRoute(lost.LostInSpace)

	routerCfg := &router.RouterConfig{
		Api: &app.RouterGroup,
		Db:  db.Db,
	}
	router.ApiRoutes(routerCfg)

	// Create an HTTP server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", appConfig.App.Port),
		Handler: app,
	}

	// Start the server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Server failed to start: %v\n", err)
		}
	}()

	gracefulShutdown(server)
}

func gracefulShutdown(server *http.Server) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	log.Println("Server is shutting down...")

	// The context is used to inform the server it has 10 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println("Server forced to shutdown:", err)
	} else {
		log.Println("Server gracefully stopped 🔴")
	}
}
