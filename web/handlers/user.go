package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/hutaochu/hello-hutao/internal/services"
)

// Register 注册
// @BasePath /api/v1
// @Summary user register
// @Tags user
// @Accept json
// @Produce json
// @Param user body services.User true "user info"
// @Success 200 {object} services.Hello
// @Router /api/v1/user/register [post]
func Register(ctx *gin.Context) {
	var user services.User
	if err := ctx.ShouldBindBodyWith(&user, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	err := services.Register(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "register successfully"})
}

// Login 登陆
func Login() {}

// Logout 退出
func Logout() {}

// Logoff 注销
func Logoff() {}

// ModifyPassword 修改密码
func ModifyPassword() {}

// ResetPassword 重制密码
func ResetPassword() {}
