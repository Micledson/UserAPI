package main

import (
	"userapi/api/controllers"

	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()

	router.GET("/users", controllers.GetUsers)
	router.POST("/users", controllers.CreateUser)

	router.Run("localhost:8080")
}