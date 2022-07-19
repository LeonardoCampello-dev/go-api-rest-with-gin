package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
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
	student := models.Student{Name: "mock-student", CPF: "41517379016", RG: "230693696"}

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

	request, _ := http.NewRequest("GET", "/mock-student", nil)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	assert.Equal(test, http.StatusOK, response.Code, "the request should return 200 if given a name parameter")

	responseMock := `{"API says":"Hi mock-student, how are you?"}`

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

	request, _ := http.NewRequest("GET", "/students/cpf/41517379016", nil)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	assert.Equal(test, http.StatusOK, response.Code)
}

func TestGetStudentById(test *testing.T) {
	database.ConnectWithDatabase()

	CreateMockStudent()

	defer DeleteMockStudent()

	router := MakeRoutes()

	router.GET("/students/:id", controllers.GetStudentById)

	request, _ := http.NewRequest("GET", "/students/"+strconv.Itoa(ID), nil)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	var mockStudent models.Student

	json.Unmarshal(response.Body.Bytes(), &mockStudent)

	assert.Equal(test, "mock-student", mockStudent.Name)
	assert.Equal(test, "41517379016", mockStudent.CPF)
	assert.Equal(test, "230693696", mockStudent.RG)
}

func TestDeleteStudentById(test *testing.T) {
	database.ConnectWithDatabase()

	CreateMockStudent()

	router := MakeRoutes()

	router.DELETE("/students/:id", controllers.DeleteStudentById)

	request, _ := http.NewRequest("DELETE", "/students/"+strconv.Itoa(ID), nil)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	assert.Equal(test, http.StatusOK, response.Code)

	if response.Code != http.StatusOK {
		DeleteMockStudent()
	}
}

func TestUpdateStudentById(test *testing.T) {
	database.ConnectWithDatabase()

	CreateMockStudent()

	defer DeleteMockStudent()

	router := MakeRoutes()

	router.PATCH("/students/:id", controllers.UpdateStudentById)

	student := models.Student{Name: "mock-student", CPF: "03224525007", RG: "280332646"}

	studentJSON, _ := json.Marshal(student)

	request, _ := http.NewRequest("PATCH", "/students/"+strconv.Itoa(ID), bytes.NewBuffer(studentJSON))

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	var updatedMockStudent models.Student

	json.Unmarshal(response.Body.Bytes(), &updatedMockStudent)

	assert.Equal(test, "mock-student", updatedMockStudent.Name)
	assert.Equal(test, "03224525007", updatedMockStudent.CPF)
	assert.Equal(test, "280332646", updatedMockStudent.RG)
}
