package middlewares

import (
	"github.com/denistort/go-booking-app/pkg/config"
	"github.com/justinas/nosurf"
	"net/http"
)

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   config.GetAppConfig().InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}
