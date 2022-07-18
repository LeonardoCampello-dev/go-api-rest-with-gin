package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LeonardoCampello-dev/go-api-rest-with-gin/controllers"
	"github.com/gin-gonic/gin"
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

	if response.Code != http.StatusOK {
		test.Fatalf("received status: %d, expected status: %d", response.Code, 200)
	}
}
