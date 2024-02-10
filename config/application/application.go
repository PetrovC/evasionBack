package application

import (
	"evasion/config/routes"
	"fmt"
	"log"
	"net/http"
)

const Version = "0.0.1"

type Configuration struct {
	Port int
	Env  string
}

type Application struct {
	Config Configuration
	Logger *log.Logger
}

func (app *Application) Configure() *http.ServeMux {
	mux := http.NewServeMux()

	routes.AddCoreRoutes(mux)

	return mux
}

func (app *Application) Init(mux *http.ServeMux) {
	addr := fmt.Sprintf(":%d", app.Config.Port)

	err := http.ListenAndServe(addr, mux)
	if err != nil {
		fmt.Println(err) // Todo handle error
	}
}
