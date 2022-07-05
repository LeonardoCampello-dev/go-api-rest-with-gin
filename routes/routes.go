package routes

import (
	"github.com/LeonardoCampello-dev/go-api-rest-with-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	router := gin.Default()

	router.GET("/:name", controllers.Salutation)
	router.GET("/students", controllers.GetAllStudents)
	router.POST("/students", controllers.CreateStudent)

	router.Run()
}
