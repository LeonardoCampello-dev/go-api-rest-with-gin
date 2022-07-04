package main

import (
	"github.com/gin-gonic/gin"
)

func GetAllStudents(context *gin.Context) {
	context.JSON(200, gin.H{
		"id":   "1",
		"name": "Leonardo",
	})
}

func main() {
	router := gin.Default()

	router.GET("/students", GetAllStudents)

	router.Run()
}
