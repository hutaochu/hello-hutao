package services

import (
	"github.com/hutaochu/hello-hutao/internal/entity"
)

type User struct {
	Name      string
	AvatarUrl string `json:"avatar_url"`
	Email     string
}

// Register 注册
func Register(user User) error {
	u := &entity.User{
		Name:      user.Name,
		AvatarUrl: user.AvatarUrl,
		Email:     user.Email,
	}
	err := entity.AddUser(u)
	return err
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

// SetPassword 设置密码
func SetPassword() {}
