package main

import (
	"github.com/denistort/go-booking-app/cmd/web/app"
	"github.com/denistort/go-booking-app/cmd/web/config"
)

func main() {
	application := app.Create().Init(config.Global)
	application.StartServer()
}
