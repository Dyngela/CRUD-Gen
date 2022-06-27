package main

import (
	"CRUD/src/api/controller"
	"CRUD/src/api/data"
	"log"
	"net/http"
)

func main() {
	db := models.GetConnection()

	models.SetDB(db)
	var appRouter = controller.CreateRouter()

	log.Println("Listening on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", appRouter))
}
