package routes

import (
	"bpjs/confiq"
	"bpjs/endpoint"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/registration", endpoint.RegistrationEndpoint).Methods("POST", "OPTIONS")
	r.HandleFunc("/login", endpoint.LoginEndpoint).Methods("POST", "OPTIONS")
	r.Handle("/user/profile", confiq.AuthMiddleware(http.HandlerFunc(endpoint.UserUpdateEndpoint))).Methods("PUT", "GET", "OPTIONS")
	r.Handle("/skills", confiq.AuthMiddleware(http.HandlerFunc(endpoint.SkillEntpointWithoutId))).Methods("POST", "GET", "OPTIONS")
	r.Handle("/skills/{Id}", confiq.AuthMiddleware(http.HandlerFunc(endpoint.SkillEntpointWithId))).Methods("PUT", "DELETE", "GET", "OPTIONS")
	r.Handle("/education", confiq.AuthMiddleware(http.HandlerFunc(endpoint.EducationEntpointWithoutId))).Methods("POST", "GET", "OPTIONS")
	r.Handle("/education/{Id}", confiq.AuthMiddleware(http.HandlerFunc(endpoint.EducationEntpointWithId))).Methods("PUT", "DELETE", "GET", "OPTIONS")

	return r
}