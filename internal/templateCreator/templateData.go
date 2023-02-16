package templateCreator

import "github.com/denistort/go-booking-app/internal/forms"

type TemplateData struct {
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}
