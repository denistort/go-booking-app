package main

import (
	"github.com/denistort/go-booking-app/cmd/web/app"
	"github.com/denistort/go-booking-app/cmd/web/config"
	"github.com/denistort/go-booking-app/cmd/web/session"
	"github.com/go-playground/validator/v10"
	"log"
)

func main() {
	Run()
}

func Run() {
	c, configErr := config.New(&config.AppConfig{
		InProduction: false,
		Port:         8080,
		UseCache:     false,
		Session:      session.Create(),
		Validator:    validator.New(),
	})
	if configErr != nil {
		log.Fatal("some troubles with config")
	}

	application, _ := app.New(c)
	startError := application.StartServer()

	if startError != nil {
		log.Fatal(startError)
	}
}
