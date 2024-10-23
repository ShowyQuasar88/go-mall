package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `gorm:"uniqueIndex;type:VARCHAR(128)"`
	Password string `gorm:"type:VARCHAR(64) not null"`
}

func (User) TableName() string {
	return "tb_user"
}