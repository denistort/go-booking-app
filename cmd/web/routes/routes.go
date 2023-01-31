package routes

import (
	"github.com/denistort/go-booking-app/cmd/web/middlewares"
	"github.com/denistort/go-booking-app/pkg/config"
	"github.com/denistort/go-booking-app/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Routes(_ *config.AppConfig) http.Handler {
	r := chi.NewRouter()
	// Middlewares
	r.Use(middlewares.Log)
	r.Use(middlewares.NoSurf)
	r.Use(middlewares.SessionLoad)
	// Paths
	r.Get("/", handlers.HomeHandler)
	r.Get("/about", handlers.AboutHandler)
	r.Get("/rooms", handlers.RoomHandler)
	r.Get("/contact-us", handlers.ContactHandler)
	r.Get("/check-available", handlers.CheckAvailableHandler)
	r.Get("/reservation", handlers.ReservationHandler)
	// Serve Static files
	fileServer := http.FileServer(http.Dir("./static/"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return r
}
