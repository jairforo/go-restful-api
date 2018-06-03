package main

import "github.com/gorilla/mux"

//Setup Routes
func SetupRouter() *mux.Router {

	router := mux.NewRouter()

	//Add prefix api/v1
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	return router
}
