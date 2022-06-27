package controller

import (
	"CRUD/src/api/services"
	"github.com/gorilla/mux"
)

var router *mux.Router

// CreateRouter -> General router who gather every sub router.
// It's him whose called in the main to initialise the server's routing/**
func CreateRouter() *mux.Router {
	router = mux.NewRouter()

	tableRoute()

	return router
}

// tableRoute -> Controller for the table /**
func tableRoute() {
	subRouterTable := router.PathPrefix("/database").Subrouter()

	subRouterTable.HandleFunc("/tables", services.FindAllTables).Methods("GET")
	subRouterTable.HandleFunc("/table/{id}", services.FindTableById).Methods("GET")
	subRouterTable.HandleFunc("/table", services.CreateTable).Methods("POST")
	subRouterTable.HandleFunc("/table", services.UpdateTable).Methods("PUT")
	subRouterTable.HandleFunc("/table/{id}", services.DeleteTableById).Methods("DELETE")
}
