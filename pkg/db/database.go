package db

import (
	"github.com/joe-l-mathew/GoFileVault/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	DB = db
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.User{})
}
