package reservation_controller

import (
	"encoding/json"
	"fmt"
	"github.com/denistort/go-booking-app/cmd/web/config"
	"github.com/denistort/go-booking-app/internal/forms"
	"github.com/denistort/go-booking-app/internal/reservation/reservation-service"
	"github.com/denistort/go-booking-app/internal/templateCreator"
	"net/http"
)

// ReservationController is handlers type
type ReservationController struct {
	AppConfig *config.AppConfig
	Templater *templateCreator.TemplateCreator
	Service   *reservation_service.ReservationService
}

// New creates the ReservationController
func New(
	config *config.AppConfig,
	t *templateCreator.TemplateCreator,
	service *reservation_service.ReservationService,
) *ReservationController {
	return &ReservationController{
		AppConfig: config,
		Templater: t,
		Service:   service,
	}
}

func (r *ReservationController) ReservationHandler(w http.ResponseWriter, req *http.Request) {
	r.Templater.Create(w, req, "reservation.page.tmpl", &templateCreator.TemplateData{
		Form: forms.New(nil),
	})
}

func (r *ReservationController) PostCreateReservation(w http.ResponseWriter, req *http.Request) {
	_ = req.ParseForm()

	//validationFields := []string{"date-from", "date-to", "phone", "email", "last_name", "first_name"}
	//mapOfBody, err := utils.ReqBodyToMap(validationFields, req)
	model := r.Service.CreateReservation(
		req.Form.Get("first_name"),
		req.Form.Get("last_name"),
		req.Form.Get("phone"),
		req.Form.Get("email"),
		req.Form.Get("date-from"),
		req.Form.Get("date-to"),
	)
	fmt.Println(model)
	resp := CreateReservationResponse{
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

type CreateReservationResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func (r *ReservationController) CheckAvailableHandler(w http.ResponseWriter, req *http.Request) {
	r.Templater.Create(w, req, "check-available.page.tmpl", &templateCreator.TemplateData{})
}
