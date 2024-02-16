package database

import (
	"github.com/anshidmattara7861/Go-Gin-backend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Initialize() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})
}
