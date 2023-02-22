package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kovalyov-valentin/go-jwt/controllers"
	"github.com/kovalyov-valentin/go-jwt/initializers"
	"github.com/kovalyov-valentin/go-jwt/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth ,controllers.Validate)


	r.Run()
}