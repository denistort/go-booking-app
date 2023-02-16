package config

import (
	"github.com/denistort/go-booking-app/cmd/web/session"
)

// Global this is our global config we can use it throughout the project
var Global = New(&AppConfig{
	InProduction: false,
	Port:         8080,
	UseCache:     false,
	Session:      session.Create(),
})
