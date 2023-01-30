package handlers

import (
	"booking-app/pkg/config"
	"booking-app/pkg/render"
	"fmt"
	"net/http"
)

// Repo is Repository used by the handlers
var Repo *Repository

// Repository is repository type
type Repository struct {
	AppConfig *config.AppConfig
}

// CreateRepo NewRepo creates the Repository
func CreateRepo(a *config.AppConfig) {
	Repo = &Repository{
		AppConfig: a,
	}
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	remoteAddr := req.RemoteAddr
	Repo.AppConfig.Session.Put(req.Context(), "remote_ip", remoteAddr)
	//fmt.Println(r.AppConfig.Session)
	render.Template(w, "home.page.tmpl", &render.TemplateData[string]{
		Data: "Something in the way",
	})
}

func AboutHandler(w http.ResponseWriter, req *http.Request) {
	remoteIp := Repo.AppConfig.Session.GetString(req.Context(), "remote_ip")
	fmt.Println(remoteIp, "REMOTE IP FROM SESSION")

	render.Template(w, "about.page.tmpl", &render.TemplateData[string]{
		Data: "Some string in the about handler",
	})
}
