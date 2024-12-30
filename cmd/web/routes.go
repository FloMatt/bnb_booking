package main

import (
	"net/http"

	"github.com/FloMatt/bnb_booking/pkg/config"
	"github.com/FloMatt/bnb_booking/pkg/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// routes initializes and configures the HTTP request multiplexer (mux) with middleware and routes.
// It returns an http.Handler that can be used to serve HTTP requests.
//
// Parameters:
// - app: A pointer to an instance of config.AppConfig, which contains application-wide configuration settings.
//
// Return value:
// - An http.Handler that can be used to serve HTTP requests.
func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/single-rooms", handlers.Repo.Single)
	mux.Get("/deluxe-rooms", handlers.Repo.Deluxe)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/reservation", handlers.Repo.Reservation)
	mux.Get("/search-availability", handlers.Repo.SearchAvailability)
	mux.Post("/search-availability", handlers.Repo.PostSearchAvailability)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
