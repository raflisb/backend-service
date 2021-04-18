package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"-"`
	Email    string `gorm:unique json:"email"`
	Token    string `json: "token" `
}

type UserDetail struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `gorm:unique json:"email"`
}
