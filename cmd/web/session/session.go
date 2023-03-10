package session

import (
	"github.com/alexedwards/scs/v2"
	"net/http"
	"time"
)

func Create() *scs.SessionManager {
	manager := scs.New()
	manager.Lifetime = 24 * time.Hour
	manager.Cookie.Persist = true
	manager.Cookie.SameSite = http.SameSiteLaxMode
	manager.Cookie.Secure = false
	return manager
}
