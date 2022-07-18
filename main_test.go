package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LeonardoCampello-dev/go-api-rest-with-gin/controllers"
	"github.com/LeonardoCampello-dev/go-api-rest-with-gin/database"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func MakeRoutes() *gin.Engine {
	routes := gin.Default()

	return routes
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

	router := MakeRoutes()

	router.GET("/students", controllers.GetAllStudents)

	request, _ := http.NewRequest("GET", "/students", nil)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	assert.Equal(test, http.StatusOK, response.Code)
}
