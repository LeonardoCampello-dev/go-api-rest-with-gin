package controllers

import (
	"github.com/LeonardoCampello-dev/go-api-rest-with-gin/models"
	"github.com/gin-gonic/gin"
)

func GetAllStudents(context *gin.Context) {
	context.JSON(200, models.Students)
}

func Salutation(context *gin.Context) {
	name := context.Params.ByName("name")

	context.JSON(200, gin.H{
		"API says": "Hi " + name + ", how are you?",
	})
}
