package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/kabilovtoha/go_study_bookings/internal/config"
	"github.com/kabilovtoha/go_study_bookings/internal/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	//mux := pat.New()
	//mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/rooms/generals-quarters", handlers.Repo.RoomsGeneralsQuarters)
	mux.Get("/rooms/majors-suite", handlers.Repo.RoomsMajorsSuite)
	mux.Get("/contacts", handlers.Repo.Contacts)
	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/api-search-availability", handlers.Repo.ApiPostAvailability)
	mux.Get("/make-reservation", handlers.Repo.MakeReservation)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
