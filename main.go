package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hutaochu/hello-hutao/docs"
	"github.com/hutaochu/hello-hutao/internal/config"
	"github.com/hutaochu/hello-hutao/internal/entity"
	"github.com/hutaochu/hello-hutao/web"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	config.InitEnv()
	config.InitConfig()
	entity.DBOpen()

	router := gin.Default()
	docs.SwaggerInfo.BasePath = ""

	v1 := router.Group("/api/v1")
	web.Routes(v1)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
		fmt.Println("server listen on :8080")
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Start shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown: \n", err)
	}

	log.Println("Server shutdown")
}
