package config

import (
	"github.com/denistort/go-booking-app/cmd/web/server"
	"github.com/denistort/go-booking-app/cmd/web/session"
	"github.com/denistort/go-booking-app/internal/config"
	"github.com/denistort/go-booking-app/internal/render"
)

var tc, _ = render.CreateTemplateCache()

// Global this is our global config we can use it throughout the project
var Global = config.Create(&config.AppConfig{
	InProduction:  false,
	Port:          server.Port,
	TemplateCache: tc,
	UseCache:      false,
	Session:       session.Create(),
})
