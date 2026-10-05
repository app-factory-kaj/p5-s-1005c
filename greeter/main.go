package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Greeting mirrors the Greeting schema in openapi.yaml.
type Greeting struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

// ErrorResponse mirrors the Error schema in openapi.yaml.
type ErrorResponse struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description,omitempty"`
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{
			Code:        http.StatusBadRequest,
			Message:     "name is required",
			Description: "the name query parameter is missing or empty",
		})
		return
	}
	writeJSON(w, http.StatusOK, Greeting{
		Name:    name,
		Message: "Hello, " + name + "!",
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", helloHandler)

	log.Println("greeter listening on :9090")
	if err := http.ListenAndServe(":9090", mux); err != nil {
		log.Fatal(err)
	}
}
