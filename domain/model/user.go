package model

import (
    "time"

	"gorm.io/gorm"
)

type Users []User

type User struct {
	ID          int64          `json:"id"`
	UserKey     string         `json:"user_key"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	Password    string         `json:"password"`
	Token       string         `json:"token"`
	Status      string         `json:"status"`
	Description string         `json:"description"`
	ImagePath   string         `json:"image_path"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
}

func EmptyUser() *User {
	return &User{}
}

func (user *User) IsEmpty() bool {
	return (
		user.ID == 0 &&
		user.UserKey == "" &&
		user.Name == "" &&
		user.Email == "" &&
		user.Password == "" &&
		user.Token == "" &&
		user.Status == "" &&
		user.Description == "")
}
