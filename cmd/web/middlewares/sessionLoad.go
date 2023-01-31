package middlewares

import (
	"github.com/denistort/go-booking-app/internal/config"
	"net/http"
)

func SessionLoad(next http.Handler) http.Handler {
	return config.GetAppConfig().Session.LoadAndSave(next)
}
