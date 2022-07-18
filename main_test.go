package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LeonardoCampello-dev/go-api-rest-with-gin/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func MakeRoutes() *gin.Engine {
	routes := gin.Default()

	return routes
}

func TestCheckSalutationStatusCode(test *testing.T) {
	routes := MakeRoutes()

	routes.GET("/:name", controllers.Salutation)

	req, _ := http.NewRequest("GET", "/fake-name", nil)

	response := httptest.NewRecorder()

	routes.ServeHTTP(response, req)

	assert.Equal(test, http.StatusOK, response.Code, "the request should return 200 if given a name parameter")

	responseMock := `{"API says":"Hi fake-name, how are you?"}`

	responseBody, _ := ioutil.ReadAll(response.Body)

	assert.Equal(test, responseMock, string(responseBody))
}
