package routes

import (
	"github.com/douglira/alura-golang-api-rest.git/controllers"
	"github.com/douglira/alura-golang-api-rest.git/middlewares"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()
	subRouterAPI := r.PathPrefix("/api").Subrouter()
	subRouterAPI.Use(middlewares.SetApplicationJsonContentType)
	subRouterAPI.HandleFunc("/personalities", controllers.AllPersonalities).Methods(http.MethodGet)
	subRouterAPI.HandleFunc("/personalities", controllers.RegisterPersonality).Methods(http.MethodPost)
	subRouterAPI.HandleFunc("/personalities/{id}", controllers.FindPersonality).Methods(http.MethodGet)
	subRouterAPI.HandleFunc("/personalities/{id}", controllers.UpdatePersonality).Methods(http.MethodPut)
	subRouterAPI.HandleFunc("/personalities/{id}", controllers.DeletePersonality).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
