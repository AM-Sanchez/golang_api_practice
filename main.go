package main

import (
	"fmt"
	"log"
	"net/http"
)

type server struct{}

func createMessage(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, message)))
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		createMessage(w, http.StatusOK, "get called")
	case "POST":
		createMessage(w, http.StatusCreated, "post called")
	case "PUT":
		createMessage(w, http.StatusAccepted, "put called")
	case "DELETE":
		createMessage(w, http.StatusOK, "delete called")
	default:
		createMessage(w, http.StatusNotFound, "not found")
	}

}

func main() {
	// hello := "hello world\n"
	// print(hello)
	s := &server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
