package main

import (
	"encoding/json"
	"net/http"

	"./types"
	"github.com/gorilla/mux"
)

//Aks to Rodrigo about it
var users []types.User

// Display all from the people var
func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

// Display a single data
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&types.User{})
}

// create a new item
func CreateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user types.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = params["id"]
	users = append(users, user)
	json.NewEncoder(w).Encode(users)
}

// Delete an item
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(users)
	}
}

// main function to boot up everything
func main() {
	users = append(users, types.User{ID: "1", Name: "John", Email: "Doe"})
	users = append(users, types.User{ID: "2", Name: "Koko", Email: "Doe"})

	router := SetupRouter()
	http.ListenAndServe(":8000", router)
}
