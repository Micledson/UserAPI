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

func GetUserByCPF(context *gin.Context) {
	cpf := context.Param("cpf")

	user := services.GetUserByCPF(cpf)

	if user == nil {
		context.IndentedJSON(404, gin.H{"message": "User not found"})
		return
	}

	context.IndentedJSON(200, user)
	
}