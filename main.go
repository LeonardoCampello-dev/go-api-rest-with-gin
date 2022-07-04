package main

import (
	"github.com/LeonardoCampello-dev/go-api-rest-with-gin/database"
	"github.com/LeonardoCampello-dev/go-api-rest-with-gin/models"
	"github.com/LeonardoCampello-dev/go-api-rest-with-gin/routes"
)

func main() {
	database.ConnectWithDatabase()

	models.Students = []models.Student{
		{Name: "Leonardo", CPF: "772.640.720-45", RG: "39.492.069-7"},
		{Name: "Guilherme", CPF: "854.077.690-18", RG: "40.856.432-5"},
	}

	routes.HandleRequests()
}
