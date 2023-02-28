package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hutaochu/hello-hutao/internal/pkg/services"
)

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
func SayHello(c *gin.Context) {
	name := c.Query("name")
	c.JSON(http.StatusOK, services.SayHello(name))
}
