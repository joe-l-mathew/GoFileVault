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
	UserId   string
	FileName string
	Filetype string
	FilePath string
}
