package main

import (
	"log"
	"net/http"
)

func main() {
	router := SetupRouter()

	log.Println("Server running on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
