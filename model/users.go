package model

import (
	"gorm.io/gorm"
)

type (
	Users struct {
		Id          int            `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
		Name        string         `gorm:"not null" json:"name"`
		SName       string         `json:"surname"`
		Username    string         `gorm:"not null" json:"username"`
		Password    string         `gorm:"not null" json:"password"`
		Email       string         `gorm:"not null" json:"email"`
		Phone       string         `gorm:"not null" json:"phone"`
		DateOfBirth string         `json:"dateofbirth"`
		Gender      string         `json:"gender"`
		Thumbnail   string         `json:"thumbnail"`
		CreatedAt   JSONTime       `gorm:"column:created_at" json:"created_at"`
		UpdatedAt   JSONTime       `gorm:"column:updated_at" json:"updated_at"`
		DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	}
)
