package render

type TemplateData[K comparable] struct {
	Data      K
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
