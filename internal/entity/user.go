package entity

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"column:id;AUTO_INCREMENT;primary_key"`        // id
	Name      string    `gorm:"column:name;NOT NULL"`                        // user name
	AvatarUrl string    `gorm:"column:avatar_url"`                           // user name
	Email     string    `gorm:"column:email;unique;NOT NULL"`                // user email
	Password  string    `gorm:"column:password;NOT NULL"`                    // user login password
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"` // create time
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"` // update time
}

func (m *User) TableName() string {
	return "user_tab"
}

func AddUser(user *User) error {
	return db.Create(user).Error
}
