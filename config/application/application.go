package application

import (
	"database/sql"
	"evasion/config/routes"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Configuration struct {
	Port    int
	Env     string
	Version string
}

type Application struct {
	Config Configuration
	Logger *log.Logger
}

func (app *Application) DadaBaseCheck() (error error) {
	db, err := Queryable()
	if err != nil {
		error = err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			error = err
		}
	}(db) // Ensure the database connection is closed when the function exits

	err = db.Ping()
	if err != nil {
		error = err
	}

	return error
}

func (app *Application) ConfigureRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	routes.AddCoreRoutes(mux)

	return mux
}

func (app *Application) Init(mux *http.ServeMux) {
	addr := fmt.Sprintf(":%d", app.Config.Port)

	// Configuration server
	srv := &http.Server{
		Addr:         addr,
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err) // Todo handle error
	}
}
