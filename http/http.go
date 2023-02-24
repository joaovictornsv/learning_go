package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func toJSON(data any) []byte {
	dataEncodedToJSONInBytes, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	return dataEncodedToJSONInBytes
}

func hello(w http.ResponseWriter, r *http.Request) {
	messageInBytes := []byte("Hello, World!")
	w.Write(messageInBytes)
}

func helloJSON(w http.ResponseWriter, r *http.Request) {
	messageInMap := map[string]string{"message": "Hello, World!"}
	messageInBytes := toJSON(messageInMap)
	w.Write(messageInBytes)
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/hellojson", helloJSON)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
