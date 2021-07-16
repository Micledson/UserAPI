package controllers

import (
	"userapi/api/Models"
	"userapi/api/services"

	"github.com/gin-gonic/gin"
)


func GetUsers(context *gin.Context) {
	context.IndentedJSON(200, services.GetUsers())
}

func CreateUser(context *gin.Context) {
	var newUser Models.User

	context.BindJSON(&newUser)

	context.IndentedJSON(201, services.CreateUser(newUser))
}