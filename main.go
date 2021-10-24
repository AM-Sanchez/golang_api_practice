package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Helper function for creating simple request messages.
func createMessage(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, message)))
}

func get(w http.ResponseWriter, r *http.Request) {
	createMessage(w, http.StatusOK, "get called")
}

func post(w http.ResponseWriter, r *http.Request) {
	createMessage(w, http.StatusCreated, "post called")
}

func put(w http.ResponseWriter, r *http.Request) {
	createMessage(w, http.StatusAccepted, "put called")
}

func delete(w http.ResponseWriter, r *http.Request) {
	createMessage(w, http.StatusOK, "delete called")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	createMessage(w, http.StatusNotFound, "not found")
}

func main() {
	hello := "Starting Server...\n"
	print(hello)
	r := mux.NewRouter() // Create gorrila mux router
	r.HandleFunc("/", get).Methods(http.MethodGet)
	r.HandleFunc("/", post).Methods(http.MethodPost)
	r.HandleFunc("/", put).Methods(http.MethodPut)
	r.HandleFunc("/", delete).Methods(http.MethodDelete)
	r.HandleFunc("/", notFound)
	log.Fatal(http.ListenAndServe(":8080", r))
}
