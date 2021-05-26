package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null" json:"name"`
	Phone    string `gorm:"type:varchar(20);not null" json:"phone"`
	Password string `gorm:"not null" json:"-"`
	Email    string `gorm:"unique;type:varchar(50);not null" json:"email"`
	Token    string `gorm:"type:varchar(100)" json: "token" `
}

type UserDetail struct {
	Address string `gorm:"type:varchar(200);" json:"address"`
	Country string `gorm:"type:varchar(50);" json:"country"`
}
