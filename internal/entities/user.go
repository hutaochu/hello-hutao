package entities

import (
	"context"
	"time"
)

type User struct {
	Id        uint      `gorm:"column:id;AUTO_INCREMENT;primary_key"`        // id
	Name      string    `gorm:"column:name;NOT NULL"`                        // name
	AvatarUrl string    `gorm:"column:avatar_url"`                           // avatar url
	Email     string    `gorm:"column:email;unique;NOT NULL"`                // user email
	Password  string    `gorm:"column:password;NOT NULL"`                    // user login password
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"` // create time
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"` // update time
}

func (m *User) TableName() string {
	return "user_tab"
}

func AddUser(ctx context.Context, user *User) error {
	return db.WithContext(ctx).Create(user).Error
}

func GetUser(ctx context.Context, userName string) (*User, error) {
	var user User
	err := db.WithContext(ctx).Find(&user).Where("name = ?", userName).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
