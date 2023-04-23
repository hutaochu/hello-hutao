package handlers

import (
	"github.com/hutaochu/hello-hutao/internal/types/req"
	"github.com/hutaochu/hello-hutao/pkg/utils"
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
	var user req.RegisterRequest
	if err := ctx.ShouldBindBodyWith(&user, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad req"})
		return
	}
	err := services.Register(ctx, user)
	if err != nil {
		utils.SetError(ctx, err)
		return
	}
	r := make(map[string]string)
	r["message"] = "register successfully"
	utils.SetResponse(ctx, r)
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
