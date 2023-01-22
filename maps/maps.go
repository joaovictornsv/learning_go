package main

import "fmt"

func main() {
	var HTTP_STATUS_CODE = map[string]int{
		"GET":          200,
		"POST":         201,
		"PUT":          204,
		"DELETE":       204,
		"NOT_FOUND":    404,
		"UNAUTHORIZED": 401,
		"FORBIDDEN":    403,
	}

	fmt.Println(HTTP_STATUS_CODE["UNAUTHORIZED"])
}
