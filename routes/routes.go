package routes

import (
	"github.com/LeonardoCampello-dev/go-api-rest-with-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	router := gin.Default()

	router.GET("/:name", controllers.Salutation)
	router.GET("/students", controllers.GetAllStudents)
	router.GET("/students/:id", controllers.GetStudentById)
	router.POST("/students", controllers.CreateStudent)
	router.PATCH("/students/:id", controllers.UpdateStudentById)
	router.DELETE("/students/:id", controllers.DeleteStudentById)

	router.Run()
}
