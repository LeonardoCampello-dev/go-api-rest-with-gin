package main

import (
	"github.com/LeonardoCampello-dev/go-api-rest-with-gin/database"
	"github.com/LeonardoCampello-dev/go-api-rest-with-gin/routes"
)

func main() {
	database.ConnectWithDatabase()

	routes.HandleRequests()
}
