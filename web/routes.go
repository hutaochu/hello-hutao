package web

import (
	"github.com/gin-gonic/gin"
	"github.com/hutaochu/hello-hutao/web/handlers"
)

func Routes(router *gin.RouterGroup) {
	router.GET("/hello", handlers.SayHello)
	router.POST("/user/register", handlers.Register)
}
