package main

import (
	"backend/internal/repository"
	"backend/internal/repository/dbrepo"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type TrackerApp struct {
	Port string
	DSN  string // Data Source Name
	DB   repository.DatabaseRepo
}

// @title TimeTracker API
// @version 1.0.0
// @description Time tracking service.

// @host localhost:8888
// @BasePath /api
func main() {
	var app TrackerApp

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	app.Port = os.Getenv("PORT")
	app.DSN = os.Getenv("DSN")

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	log.Printf("Listening on port %s", app.Port)
	log.Fatal(http.ListenAndServe(":"+app.Port, app.Routes()))
}
