package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type TrackerApp struct {
	Port string
	DSN  string // Data Source Name
}

func main() {
	var app TrackerApp

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	app.Port = os.Getenv("PORT")
	app.DSN = os.Getenv("DSN")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Hello, World!")
	})

	log.Printf("Listening on port %s", app.Port)
	log.Fatal(http.ListenAndServe(":"+app.Port, nil))
}
