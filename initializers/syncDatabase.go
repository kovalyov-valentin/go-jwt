package initializers

import "github.com/kovalyov-valentin/go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}