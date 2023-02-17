package user

import (
	"github.com/denistort/go-booking-app/cmd/web/config"
	reservation_service "github.com/denistort/go-booking-app/internal/reservation/reservation-service"
	"github.com/denistort/go-booking-app/internal/templateCreator"
	userController "github.com/denistort/go-booking-app/internal/user/user-controller"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func MakeUserRouter(config *config.AppConfig, t *templateCreator.TemplateCreator) func() http.Handler {
	service := reservation_service.New()
	controller := userController.New(config, t, service)

	return func() http.Handler {
		r := chi.NewRouter()
		r.Post("/sign-in", controller.SignIn)
		r.Post("/sign-up", controller.SignUp)
		return r
	}
}
