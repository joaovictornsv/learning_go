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

	fmt.Println("Server is running on port 5000!")
	http.ListenAndServe(":5000", router)
}
