package app

import (
	"fmt"
	"github.com/denistort/go-booking-app/cmd/web/appRouter"
	"github.com/denistort/go-booking-app/cmd/web/config"
	"log"
	"net/http"
	"strconv"
)

var created = false

type app struct {
	server *http.Server
}

func New(config *config.AppConfig) *app {
	if created == true {
		panic("App cant be created two times")
	} else {
		created = true
		app := &app{}
		router := appRouter.New(config)
		fmt.Println(":" + strconv.Itoa(config.Port))
		app.server = &http.Server{
			Addr:    ":" + strconv.Itoa(config.Port),
			Handler: router.Start(),
		}
		return app
	}
}

func (app *app) StartServer() {
	startErr := app.server.ListenAndServe()
	if startErr != nil {
		log.Fatal("Server isn't Started")
	}
}
