package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hutaochu/hello-hutao/constants"
	"github.com/hutaochu/hello-hutao/docs"
	"github.com/hutaochu/hello-hutao/internal/config"
	"github.com/hutaochu/hello-hutao/internal/entities"
	"github.com/hutaochu/hello-hutao/pkg/middlewares"
	"github.com/hutaochu/hello-hutao/pkg/utils/logger"
	"github.com/hutaochu/hello-hutao/web"
	"github.com/hutaochu/hello-hutao/web/handlers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config.InitEnv()
	config.InitConfig()
	entities.DBOpen()
	router := gin.Default()
	docs.SwaggerInfo.BasePath = ""

	port := os.Getenv("PORT")
	if port == "" {
		port = constants.DefaultPort
	}

	v1 := router.Group("/apis/v1")
	v1.Use(middlewares.Trace())
	v1.Use(middlewares.RequestLogger())
	v1.Use(middlewares.HandleResponse())

	web.Routes(v1)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/public_key", middlewares.HandleResponse(), handlers.GetPublicKey)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	logger.Logger.Info(fmt.Sprintf("server start to listen on: %s", port))
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	s := <-quit

	logger.Logger.Info(fmt.Sprintf("Start shutdown server with signal: %s", s.String()))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown: \n", err)
	}

	logger.Logger.Info("Server shutdown")
}
