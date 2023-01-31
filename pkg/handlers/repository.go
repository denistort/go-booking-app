package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/denistort/go-booking-app/pkg/config"
	"github.com/denistort/go-booking-app/pkg/render"
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
	render.Template(w, req, "home.page.tmpl", &render.TemplateData[string]{
		Data: "Something in the way",
	})
}

func AboutHandler(w http.ResponseWriter, req *http.Request) {
	remoteIp := Repo.AppConfig.Session.GetString(req.Context(), "remote_ip")
	fmt.Println(remoteIp, "REMOTE IP FROM SESSION")

	render.Template(w, req, "about.page.tmpl", &render.TemplateData[string]{
		Data: "Some string in the about handler",
	})
}
func RoomHandler(w http.ResponseWriter, req *http.Request) {
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
	render.Template(w, req, templateName, &render.TemplateData[string]{
		Data: "Some string in the about handler",
	})
}

func ContactHandler(w http.ResponseWriter, req *http.Request) {
	render.Template(w, req, "contact.page.tmpl", &render.TemplateData[string]{
		Data: "Some string in the about handler",
	})
}

func CheckAvailableHandler(w http.ResponseWriter, req *http.Request) {
	render.Template(w, req, "check-available.page.tmpl", &render.TemplateData[string]{
		Data: "Some string in the about handler",
	})
}

func ReservationHandler(w http.ResponseWriter, req *http.Request) {
	render.Template(w, req, "reservation.page.tmpl", &render.TemplateData[string]{
		Data: "Some string in the about handler",
	})
}

type JsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func PostReservationHandler(w http.ResponseWriter, req *http.Request) {
	dateFrom := req.Form.Get("date-from")
	dateTo := req.Form.Get("date-to")
	fmt.Println(dateFrom, dateTo)
	resp := JsonResponse{
		Ok:      true,
		Message: "Я работаю друг все хорошо",
	}
	out, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
