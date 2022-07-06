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
	var students []models.Student

	database.DB.Find(&students)

	context.JSON(http.StatusOK, students)
}

func GetStudentById(context *gin.Context) {
	var student models.Student

	id := context.Params.ByName("id")

	database.DB.First(&student, id)

	if student.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "student not found",
		})

		return
	}

	context.JSON(http.StatusOK, student)
}

func GetStudentByCPF(context *gin.Context) {
	var student models.Student

	cpf := context.Params.ByName("cpf")

	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "student not found",
		})

		return
	}

	context.JSON(http.StatusOK, student)
}

func CreateStudent(context *gin.Context) {
	var student models.Student

	if err := context.ShouldBindJSON(&student); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	if err := models.ValidateStudentSchema(&student); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	database.DB.Create(&student)

	context.JSON(http.StatusOK, student)
}

func UpdateStudentById(context *gin.Context) {
	var student models.Student

	id := context.Params.ByName("id")

	database.DB.First(&student, id)

	if err := context.ShouldBindJSON(&student); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	if err := models.ValidateStudentSchema(&student); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	database.DB.Model(&student).UpdateColumns(student)

	context.JSON(http.StatusOK, student)
}

func DeleteStudentById(context *gin.Context) {
	var student models.Student

	id := context.Params.ByName("id")

	database.DB.Delete(&student, id)

	context.JSON(http.StatusOK, gin.H{
		"message": "student successfully deleted",
	})
}
