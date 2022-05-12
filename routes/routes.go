package routes

import (
	"github.com/douglira/alura-golang-api-rest.git/controllers"
	"github.com/gorilla/mux"

	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()
	subRouterAPI := r.PathPrefix("/api").Subrouter()
	subRouterAPI.HandleFunc("/personalities", controllers.AllPersonalities).Methods("GET")
	subRouterAPI.HandleFunc("/personalities/{id}", controllers.FindPersonality).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
