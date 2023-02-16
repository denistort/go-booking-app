package middlewares

import (
	"github.com/denistort/go-booking-app/cmd/web/config"
	"net/http"
)

func SessionLoad(next http.Handler) http.Handler {
	return config.Global.Session.LoadAndSave(next)
}
