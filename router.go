package main

import (
	"encoding/json"
	"net/http"
)

func SetupRouter() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/health", LoggingMiddleware(http.HandlerFunc(healthHandler)))
	mux.Handle("/hello", LoggingMiddleware(http.HandlerFunc(helloHandler)))

	return mux
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"message": "Hello from Go server!"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
