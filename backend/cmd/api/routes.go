package main

import (
	_ "backend/docs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func (app *TrackerApp) Routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(zapLogger)
	mux.Use(enableCors)

	mux.Route("/api", func(mux chi.Router) {
		mux.Route("/users", func(mux chi.Router) {
			mux.Get("/", app.GetAllUsers)
			mux.Put("/", app.AddUser)
			mux.Post("/", app.GetUserByPassport)
			mux.Post("/{id}", nil)
			mux.Patch("/{id}", nil)
			mux.Delete("/{id}", nil)
		})
	})

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello World"))
	})

	mux.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+app.Port+"/swagger/doc.json"),
	))

	return mux
}
