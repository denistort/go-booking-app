package routes

import (
	"github.com/denistort/go-booking-app/cmd/web/middlewares"
	"github.com/denistort/go-booking-app/pkg/config"
	"github.com/denistort/go-booking-app/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Routes(config *config.AppConfig) http.Handler {
	r := chi.NewRouter()
	// Middlewares
	r.Use(middlewares.Log)
	r.Use(middlewares.NoSurf)
	r.Use(middlewares.SessionLoad)
	// Paths
	r.Get("/", handlers.HomeHandler)
	r.Get("/about", handlers.AboutHandler)
	return r
}
