package web

import (
	"github.com/gin-gonic/gin"
	"github.com/hutaochu/hello-hutao/docs"
	"github.com/hutaochu/hello-hutao/web/handlers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RunApp() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = ""

	v1 := router.Group("/api/v1")
	v1.GET("/hello", handlers.SayHello)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
