package model

import (
	"time"

	"gorm.io/gorm"
)

type DirectChat struct {
	ID              int64     `json:"id"`
	DirectChatKey   string    `json:"direct_mail_key"`
	MutualFollowKey string    `json:"mutual_follow_key"`
	UserKey         string    `json:"user_key"`
	UserName        string    `json:"user_name"`
	Content         string    `json:"content"`
	ImagePath       string    `json:"image_path"`
	PostedAt        time.Time `json:"posted_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func EmptyDirectMail() *DirectChat {
	return &DirectChat{}
}

func (directMail *DirectChat) IsEmpty() bool {
	return (
		directMail.ID == 0 &&
		directMail.DirectChatKey == "" &&
		directMail.MutualFollowKey == "" &&
		directMail.UserKey == "" &&
		directMail.UserName == "" &&
		directMail.Content == "" &&
		directMail.ImagePath == "" &&
		directMail.PostedAt.IsZero())
}
