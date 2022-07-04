package main

import (
	"github.com/LeonardoCampello-dev/go-api-rest-with-gin/models"
	"github.com/LeonardoCampello-dev/go-api-rest-with-gin/routes"
)

func main() {
	models.Students = []models.Student{
		{Id: 1, Name: "Leonardo", CPF: "772.640.720-45", RG: "39.492.069-7"},
		{Id: 2, Name: "Guilherme", CPF: "854.077.690-18", RG: "40.856.432-5"},
	}

	routes.HandleRequests()
}
