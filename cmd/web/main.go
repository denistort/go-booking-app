package main

import (
	"booking-app/cmd/web/app"
	"booking-app/cmd/web/server"
	"booking-app/cmd/web/session"
	"booking-app/pkg/config"
	"booking-app/pkg/render"
)

func main() {
	tc, _ := render.CreateTemplateCache()

	appConfig := config.Create(&config.AppConfig{
		InProduction:  false,
		Port:          server.Port,
		TemplateCache: tc,
		UseCache:      false,
		Session:       session.Create(),
	})

	application := app.Create().Init(appConfig)
	application.StartServer()
}
