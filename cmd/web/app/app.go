package app

import (
	"fmt"
	"github.com/denistort/go-booking-app/cmd/web/routes"
	"github.com/denistort/go-booking-app/pkg/config"
	"github.com/denistort/go-booking-app/pkg/handlers"
	"log"
	"net/http"
)

var created = false

type app struct {
	server    *http.Server
	appConfig *config.AppConfig
}

func Create() *app {
	if created == true {
		panic("App cant be created two times")
	} else {
		created = true
		app := &app{}
		return app
	}
}

func (app *app) Init(options *config.AppConfig) *app {
	app.appConfig = options
	app.server = &http.Server{
		Addr:    options.Port,
		Handler: routes.Routes(app.appConfig),
	}
	// repo creates
	handlers.CreateRepo(app.appConfig)

	return app
}
func (app *app) StartServer() {
	startErr := app.server.ListenAndServe()

	if startErr != nil {
		log.Fatal("Server isn't Started")
	} else {
		fmt.Println("Server started on ", app.appConfig.Port)
	}
}