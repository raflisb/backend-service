package model

import (
	"gorm.io/gorm"
)

type User struct { 
	gorm.Model 
	Name		string 	`json:"name"`
	Phone		string	`json:"phone"`
	Password	string	`json:"-"`
	Email		string	`gorm:unique json:"email"` 

}