package app

import (
	"errors"
	"github.com/denistort/go-booking-app/cmd/web/appRouter"
	"github.com/denistort/go-booking-app/cmd/web/config"
	"net/http"
	"strconv"
)

var (
	AlreadyCreatedError = errors.New("app already created")
)

var created = false

type app struct {
	server *http.Server
}

func New(config *config.AppConfig) (*app, error) {
	if created == true {
		return nil, AlreadyCreatedError
	}
	created = true
	router := appRouter.New(config)
	port := ":" + strconv.Itoa(config.Port)
	app := &app{
		server: &http.Server{
			Addr:    port,
			Handler: router.Start(),
		},
	}
	return app, nil
}

func (app *app) StartServer() error {
	return app.server.ListenAndServe()
}
