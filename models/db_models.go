package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Password string `json:"-"`
	Email    string `gorm:"unique"`
}

type UserFiles struct {
	gorm.Model
	FileName string
	Filetype string
}
