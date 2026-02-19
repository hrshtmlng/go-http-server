package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := SetupRouter()

	log.Println("Server running on :" + port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
