package main

import (
	"evasion/config/application"
	"flag"
	"log"
	"os"
)

func main() {
	var app application.Application
	var cfg application.Configuration

	flag.IntVar(&cfg.Port, "port", 8080, "API server port")
	flag.StringVar(&cfg.Env, "env", "dev", "Environment (dev|test|prod)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app = application.Application{
		Config: cfg,
		Logger: logger,
	}

	mux := app.Configure()

	app.Init(mux)
}
