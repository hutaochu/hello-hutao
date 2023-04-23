package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hutaochu/hello-hutao/internal/services"
	"github.com/hutaochu/hello-hutao/pkg/utils"
)

// SayHello
// @BasePath /api/v1
// @Summary get hello
// @Schemes
// @Description get hello
// @Tags hello-hutao
// @Accept json
// @Produce json
// @Param name query string false "name"
// @Success 200 {object} services.Hello
// @Router /api/v1/hello [get]
func SayHello(ctx *gin.Context) {
	name := ctx.Query("name")
	r := services.SayHello(name)
	utils.SetResponse(ctx, r)
}
