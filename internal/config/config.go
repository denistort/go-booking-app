package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

// AppConfig holds the application configuration
type AppConfig struct {
	InProduction  bool
	Port          string
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	Session       *scs.SessionManager
}

var appConfig AppConfig

func Create(config *AppConfig) *AppConfig {
	appConfig = *config
	return &appConfig
}

func GetAppConfig() *AppConfig {
	return &appConfig
}
