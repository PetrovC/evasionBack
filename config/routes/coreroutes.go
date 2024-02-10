package routes

import (
	"evasion/config/handlers"
	"net/http"
)

func AddCoreRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/Authenticate", handlers.Authenticate)
}
