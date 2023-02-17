package mock

import (
	"github.com/denistort/go-booking-app/cmd/web/config"
	"github.com/denistort/go-booking-app/cmd/web/session"
	"github.com/go-playground/validator/v10"
)

var Config = &config.AppConfig{
	InProduction: false,
	Port:         8080,
	UseCache:     false,
	Session:      session.Create(),
	Validator:    validator.New(),
}
