package services

import (
	"context"
	"errors"
	"github.com/hutaochu/hello-hutao/internal/entities"
	"github.com/hutaochu/hello-hutao/internal/types/req"
	"gorm.io/gorm"
)

type User struct {
	Name      string
	AvatarUrl string `json:"avatar_url"`
	Email     string
}

// Register 注册
func Register(ctx context.Context, user req.RegisterRequest) error {
	u, err := entities.GetUser(ctx, user.Name)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if u != nil && u.Name != "" {
		// username is invalid
		return errors.New("duplicate username")
	}
	// 处理密码
	newUser := &entities.User{
		Name:      user.Name,
		AvatarUrl: user.AvatarUrl,
		Email:     user.Email,
		Password:  user.Password,
	}
	err = entities.AddUser(ctx, newUser)
	return err
}

// Login 登陆
func Login(email, password string) {}

// Logout 退出
func Logout() {}

// Logoff 注销
func Logoff() {}

// ModifyPassword 修改密码
func ModifyPassword(userId int, new, old string) {}

// ResetPassword 重制密码
func ResetPassword(userId int) {}

// SetPassword 设置密码
func SetPassword(userId int, passsword string) {

}
