package controllers

import (
	"userapi/api/services"

	"github.com/gin-gonic/gin"
)


func GetUsers(context *gin.Context) {
	context.IndentedJSON(200, services.GetUsers())
}