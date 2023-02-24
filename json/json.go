package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func toJSON(data any) *bytes.Buffer {
	dataEncodedToJSONInBytes, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	dataEncodedToJSON := bytes.NewBuffer(dataEncodedToJSONInBytes)

	return dataEncodedToJSON
}

func toMap(bytesArray []byte, mapReference *map[string]string) {
	err := json.Unmarshal(bytesArray, mapReference)

	if err != nil {
		log.Fatal(err)
	}
}

func toStruct(bytesArray []byte, mapReference *Task) {
	err := json.Unmarshal(bytesArray, mapReference)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// STRUCT to JSON
	taskStruct := Task{"Beber Água", "2L todos os dias"}
	taskStructInJSON := toJSON(taskStruct)

	fmt.Println("task (struct to json):", taskStructInJSON)

	// JSON to STRUCT
	var taskStructFromJSON Task
	jsonToStruct := `{"title":"Beber Água","description":"2L todos os dias"}`
	jsonBytesToStruct := []byte(jsonToStruct)
	toStruct(jsonBytesToStruct, &taskStructFromJSON)

	fmt.Println("task (json to struct):", taskStructFromJSON)

	// MAP to JSON
	taskMap := map[string]string{
		"title":       "Beber Água",
		"description": "2L todos os dias",
	}
	taskMapInJSON := toJSON(taskMap)

	fmt.Println("task (map to json):", taskMapInJSON)

	// JSON to MAP
	var taskMapFromJSON map[string]string
	jsonToMap := `{"description":"2L todos os dias","title":"Beber Água"}`
	jsonBytesToMap := []byte(jsonToMap)
	toMap(jsonBytesToMap, &taskMapFromJSON)

	fmt.Println("task (json to map):", taskMapFromJSON)
}
