package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LeonardoCampello-dev/go-api-rest-with-gin/controllers"
	"github.com/LeonardoCampello-dev/go-api-rest-with-gin/database"
	"github.com/LeonardoCampello-dev/go-api-rest-with-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func MakeRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	routes := gin.Default()

	return routes
}

var ID int

func CreateMockStudent() {
	student := models.Student{Name: "fake-student", CPF: "415.173.790-16", RG: "23.069.369-6"}

	database.DB.Create(&student)

	ID = int(student.ID)
}

func DeleteMockStudent() {
	var student models.Student

	database.DB.Delete(&student, ID)
}

func TestCheckSalutationStatusCode(test *testing.T) {
	router := MakeRoutes()

	router.GET("/:name", controllers.Salutation)

	request, _ := http.NewRequest("GET", "/fake-name", nil)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	assert.Equal(test, http.StatusOK, response.Code, "the request should return 200 if given a name parameter")

	responseMock := `{"API says":"Hi fake-name, how are you?"}`

	responseBody, _ := ioutil.ReadAll(response.Body)

	assert.Equal(test, responseMock, string(responseBody))
}

func TestStudentList(test *testing.T) {
	database.ConnectWithDatabase()

	CreateMockStudent()

	defer DeleteMockStudent()

	router := MakeRoutes()

	router.GET("/students", controllers.GetAllStudents)

	request, _ := http.NewRequest("GET", "/students", nil)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	assert.Equal(test, http.StatusOK, response.Code)
}

func TestStudentSearchByCPF(test *testing.T) {
	database.ConnectWithDatabase()

	CreateMockStudent()

	defer DeleteMockStudent()

	router := MakeRoutes()

	router.GET("/students/cpf/:cpf", controllers.GetStudentByCPF)

	request, _ := http.NewRequest("GET", "/students/cpf/415.173.790-16", nil)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	assert.Equal(test, http.StatusOK, response.Code)
}
