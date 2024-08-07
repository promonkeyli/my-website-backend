package model

import (
	"gorm.io/gorm"
	"web_backend.com/m/v2/internal/app/repository"
)

// User represents a tool entity
// @Description User represents a user entity
type User struct {
	ID       uint   `json:"id,omitempty" gorm:"column:id" swaggerignore:"true"`
	UserName string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
}

// TableName 自定义表名
func (User) TableName() string {
	return "user"
}

func QueryUser(username string) (*gorm.DB, User) {
	var user User
	return repository.DB.Where("username = ?", username).First(&user), user
}
