package routes

import (
	"evasion/config/handlers"
	"net/http"
)

func AddCoreRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/users/login", handlers.UserLogin)
	mux.HandleFunc("/users/signup", handlers.UserSignup)
	mux.HandleFunc("/api/healthcheck", handlers.HealthCheck)
}
