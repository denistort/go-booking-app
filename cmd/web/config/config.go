package config

import (
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/validator/v10"
	"html/template"
	"log"
)

var appConfig AppConfig

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

func New(config *AppConfig) *AppConfig {
	appConfig = *config
	return &appConfig
}
