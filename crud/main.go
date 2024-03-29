package main

import (
	"crud/controllers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", controllers.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", controllers.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", controllers.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods(http.MethodDelete)

	fmt.Println("Server is running on port 5000!")
	http.ListenAndServe(":5000", router)
}
