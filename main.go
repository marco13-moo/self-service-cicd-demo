package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type response struct {
	Service string `json:"service"`
	Status  string `json:"status"`
	Version string `json:"version"`
}

func main() {
	version := os.Getenv("SERVICE_VERSION")
	if version == "" {
		version = "local"
	}

	handler := http.NewServeMux()
	handler.HandleFunc("GET /", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response{Service: "self-service-cicd-demo", Status: "ok", Version: version})
	})
	handler.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok\n"))
	})

	address := ":8080"
	log.Printf("self-service-cicd-demo listening on %s", address)
	log.Fatal(http.ListenAndServe(address, handler))
}
