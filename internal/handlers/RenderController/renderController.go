package RenderController

import (
	"fmt"
	"github.com/denistort/go-booking-app/cmd/web/config"
	"github.com/denistort/go-booking-app/internal/forms"
	"github.com/denistort/go-booking-app/internal/templateCreator"
	"net/http"
)

// RenderController is handlers type
type RenderController struct {
	AppConfig *config.AppConfig
	Templater *templateCreator.TemplateCreator
}

// New creates the RenderController
func New(config *config.AppConfig, t *templateCreator.TemplateCreator) *RenderController {
	return &RenderController{
		Templater: t,
		AppConfig: config,
	}
}

func (r *RenderController) HomeHandler(w http.ResponseWriter, req *http.Request) {
	remoteAddr := req.RemoteAddr
	r.AppConfig.Session.Put(req.Context(), "remote_ip", remoteAddr)

	r.Templater.Create(w, req, "home.page.tmpl", &templateCreator.TemplateData{})
}

func (r *RenderController) AboutHandler(w http.ResponseWriter, req *http.Request) {
	remoteIp := r.AppConfig.Session.GetString(req.Context(), "remote_ip")
	fmt.Println(remoteIp, "REMOTE IP FROM SESSION")

	r.Templater.Create(w, req, "about.page.tmpl", &templateCreator.TemplateData{})
}
func (r *RenderController) RoomHandler(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	t, ok := query["t"]
	if ok != true {
		return
	}
	var templateName string
	switch t[0] {
	case "generals-quarters":
		templateName = "general.page.tmpl"
	case "majors-suite":
		templateName = "major.page.tmpl"
	default:
		templateName = "home.page.tmpl"
	}
	r.Templater.Create(w, req, templateName, &templateCreator.TemplateData{})
}

func (r *RenderController) ContactHandler(w http.ResponseWriter, req *http.Request) {
	r.Templater.Create(w, req, "contact.page.tmpl", &templateCreator.TemplateData{})
}

func (r *RenderController) CheckAvailableHandler(w http.ResponseWriter, req *http.Request) {
	r.Templater.Create(w, req, "check-available.page.tmpl", &templateCreator.TemplateData{})
}

func (r *RenderController) ReservationHandler(w http.ResponseWriter, req *http.Request) {
	r.Templater.Create(w, req, "reservation.page.tmpl", &templateCreator.TemplateData{
		Form: forms.New(nil),
	})
}
