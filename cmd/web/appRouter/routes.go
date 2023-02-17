package appRouter

import (
	"github.com/denistort/go-booking-app/cmd/web/config"
	"github.com/denistort/go-booking-app/cmd/web/middlewares"
	"github.com/denistort/go-booking-app/internal/handlers/JsonController"
	"github.com/denistort/go-booking-app/internal/handlers/RenderController"
	"github.com/denistort/go-booking-app/internal/reservation"
	"github.com/denistort/go-booking-app/internal/templateCreator"
	"github.com/denistort/go-booking-app/internal/user"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type AppRouter struct {
	router    chi.Router
	appConfig *config.AppConfig
}

func New(config *config.AppConfig) *AppRouter {
	return &AppRouter{
		router:    chi.NewRouter(),
		appConfig: config,
	}
}

func (r *AppRouter) Start() http.Handler {
	// Middlewares
	m := middlewares.New(r.appConfig)
	r.router.Use(m.Log)
	r.router.Use(m.NoSurf)
	r.router.Use(m.SessionLoad)

	// Creating Templater
	t := templateCreator.New(r.appConfig)
	// Controllers creation
	renderController := RenderController.New(r.appConfig, t)
	jsonController := JsonController.New(r.appConfig)
	reservationRouter := reservation.MakeReservationRouter(r.appConfig, t)
	userRouter := user.MakeUserRouter(r.appConfig, t)

	// Handler that templateCreator templates and getting back as response
	r.router.Get("/", renderController.HomeHandler)
	r.router.Get("/about", renderController.AboutHandler)
	r.router.Get("/rooms", renderController.RoomHandler)
	r.router.Get("/contact-us", renderController.HomeHandler)
	r.router.Get("/check-available", renderController.CheckAvailableHandler)
	// Reservation Controller Mount
	r.router.Mount("/reservation", reservationRouter())
	r.router.Mount("/", userRouter())
	// Json Handlers
	r.router.Post("/contact-us", jsonController.ContactJsonHandler)
	// Serve Static files
	fileServer := http.FileServer(http.Dir("./static/"))
	r.router.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return r.router
}
