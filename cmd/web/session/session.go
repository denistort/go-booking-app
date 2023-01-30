package session

import (
	"github.com/alexedwards/scs/v2"
	"github.com/denistort/go-booking-app/pkg/config"
	"net/http"
	"time"
)

func Create() *scs.SessionManager {
	manager := scs.New()
	manager.Lifetime = 24 * time.Hour
	manager.Cookie.Persist = true
	manager.Cookie.SameSite = http.SameSiteLaxMode
	manager.Cookie.Secure = config.GetAppConfig().InProduction
	return manager
}
