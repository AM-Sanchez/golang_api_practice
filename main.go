// Golang API
//
// V1: Requests sent to localhost:8080/api/v1/
// will get responses indicating the type of request received.
// e.g. POST localhost:8080/api/v1 will return json message of {"message": "post called"}
//
// In addition, GET requests for /user/{userID}/comment/{commentID}, where both parameters are ints,
// are handled in addition to supporting a query for "location" at this resource .
//
// Example GET request: GET localhost:8080/api/v1/user/42/comment/39?location=BlogAtTheEndOfTheUniverse

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv" //Atoi

	"github.com/gorilla/mux"
)

const gUserIDField = "userID"
const gCommentIDField = "commentID"

// Helper function for creating simple request messages.
func createMessage(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, message)))
}

func fromHomePath(r *http.Request) string {
	fromHome := ""
	if r.URL.Path == "/" {
		fromHome = " from home path"
	}
	return fromHome
}

func get(w http.ResponseWriter, r *http.Request) {
	createMessage(w, http.StatusOK, fmt.Sprintf("get called%s", fromHomePath(r)))
}

func post(w http.ResponseWriter, r *http.Request) {
	createMessage(w, http.StatusCreated, fmt.Sprintf("post called%s", fromHomePath(r)))
}

func put(w http.ResponseWriter, r *http.Request) {
	createMessage(w, http.StatusAccepted, fmt.Sprintf("put called%s", fromHomePath(r)))
}

func delete(w http.ResponseWriter, r *http.Request) {
	createMessage(w, http.StatusOK, fmt.Sprintf("delete called%s", fromHomePath(r)))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	createMessage(w, http.StatusNotFound, fmt.Sprintf("request not supported (%s)", fromHomePath(r)))
}

func params(w http.ResponseWriter, r *http.Request) {

	pathParams := mux.Vars(r) // map for key/value pairs from {key} to value as seen in the url path

	w.Header().Set("Content-Type", "appliation/json")

	// Handle param for {userID} field.
	userID := -1
	var err error
	if val, ok := pathParams[gUserIDField]; ok {
		userID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number"}`))
			return
		}
	}

	// Handle param for {commentID} field.
	commentID := -1
	if val, ok := pathParams[gCommentIDField]; ok {
		commentID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number"}`))
			return
		}
	}

	query := r.URL.Query()
	// println(fmt.Sprintf("Query: %s", query))
	location := query.Get("location")

	w.Write([]byte(fmt.Sprintf(`{"%s": %d, "%s": %d, "location": "%s" }`, gUserIDField, userID, gCommentIDField, commentID, location)))
}

func main() {
	hello := "Starting v1 Server...\n"
	print(hello)

	mainRouter := mux.NewRouter().StrictSlash(true) // Create gorrila mux router

	// Handlers for "/" resource.
	mainRouter.HandleFunc("/", get).Methods(http.MethodGet)
	mainRouter.HandleFunc("/", post).Methods(http.MethodPost)
	mainRouter.HandleFunc("/", put).Methods(http.MethodPut)
	mainRouter.HandleFunc("/", delete).Methods(http.MethodDelete)

	// Create subroute for "/api" resource.
	api := mainRouter.PathPrefix("/api").Subrouter()
	// Create subroute for "/api/v1" resource.
	apiV1 := api.PathPrefix("/v1").Subrouter()
	apiV1.HandleFunc("", get).Methods(http.MethodGet)
	apiV1.HandleFunc("", post).Methods(http.MethodPost)
	apiV1.HandleFunc("", put).Methods(http.MethodPut)
	apiV1.HandleFunc("", delete).Methods(http.MethodDelete)
	apiV1.HandleFunc("", notFound)

	// Handles GET requests for /user/{userID}/comment/{commentID}
	// The handler params() also handles queries at this resource via location.
	// Example GET request: GET localhost:8080/api/v1/user/42/comment/39?location=BlogAtTheEndOfTheUniverse
	apiV1.HandleFunc(fmt.Sprintf("/user/{%s}/comment/{%s}", gUserIDField, gCommentIDField), params).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", mainRouter))
}
