package middlewares

import "github.com/denistort/go-booking-app/cmd/web/config"

type Middlewares struct {
	Config *config.AppConfig
}

func New(c *config.AppConfig) *Middlewares {
	return &Middlewares{
		Config: c,
	}
}
