package config

import (
	"errors"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/validator/v10"
	"html/template"
	"log"
)

var AlreadyCreatedError = errors.New("already created config")
var appConfig AppConfig
var created = false

// AppConfig holds the application configuration
type AppConfig struct {
	InProduction  bool
	Port          int
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	Session       *scs.SessionManager
	Validator     *validator.Validate
}

func New(config *AppConfig) (*AppConfig, error) {
	if created == true {
		return nil, AlreadyCreatedError
	}
	appConfig = *config
	created = true
	return &appConfig, nil
}
