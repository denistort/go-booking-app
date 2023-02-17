package middlewares

import (
	"net/http"
)

func (m *Middlewares) SessionLoad(next http.Handler) http.Handler {
	return m.Config.Session.LoadAndSave(next)
}
