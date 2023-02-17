package reservation

import (
	"github.com/denistort/go-booking-app/cmd/web/config"
	"github.com/denistort/go-booking-app/internal/reservation/reservation-controller"
	"github.com/denistort/go-booking-app/internal/reservation/reservation-service"
	"github.com/denistort/go-booking-app/internal/templateCreator"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func MakeReservationRouter(config *config.AppConfig, t *templateCreator.TemplateCreator) func() http.Handler {
	service := reservation_service.New()
	controller := reservation_controller.New(config, t, service)

	return func() http.Handler {
		r := chi.NewRouter()
		r.Get("/", controller.ReservationHandler)
		r.Get("/check-available", controller.CheckAvailableHandler)
		r.Post("/", controller.PostCreateReservation)
		return r
	}
}
