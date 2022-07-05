package controllers

import (
	"net/http"

	"github.com/LeonardoCampello-dev/go-api-rest-with-gin/database"
	"github.com/LeonardoCampello-dev/go-api-rest-with-gin/models"
	"github.com/gin-gonic/gin"
)

func Salutation(context *gin.Context) {
	name := context.Params.ByName("name")

	context.JSON(http.StatusOK, gin.H{
		"API says": "Hi " + name + ", how are you?",
	})
}

func GetAllStudents(context *gin.Context) {
	context.JSON(http.StatusOK, models.Students)
}

func CreateStudent(context *gin.Context) {
	var student models.Student

	err := context.ShouldBindJSON(&student)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	database.DB.Create(&student)

	context.JSON(http.StatusOK, student)
}
