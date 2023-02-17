package reservation_controller

import (
	"github.com/denistort/go-booking-app/cmd/web/config"
	"github.com/denistort/go-booking-app/internal/forms"
	reservationdto "github.com/denistort/go-booking-app/internal/reservation/reservation-dto"
	"github.com/denistort/go-booking-app/internal/reservation/reservation-service"
	"github.com/denistort/go-booking-app/internal/templateCreator"
	"github.com/denistort/go-booking-app/internal/utils"
	"github.com/go-playground/validator/v10"
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
	reservationDto := reservationdto.ReservationDto{
		FirstName: req.Form.Get("first_name"),
		LastName:  req.Form.Get("last_name"),
		Phone:     req.Form.Get("phone"),
		Email:     req.Form.Get("email"),
		DateFrom:  req.Form.Get("date-from"),
		DateTo:    req.Form.Get("date-to"),
	}
	if validationsErrors := r.AppConfig.Validator.Struct(reservationDto); validationsErrors != nil {
		errorResponse := ValidationErrorResponse{
			Ok:      false,
			Message: "This fields is uncorrected",
		}
		for _, err := range validationsErrors.(validator.ValidationErrors) {
			errorResponse.NotValidFields = append(errorResponse.NotValidFields, Field{
				Name:  err.Field(),
				Value: err.Param(),
			})
		}
		utils.SendJson(errorResponse, w)
		return
	}
	utils.SendJson(CreateReservationResponse{
		Ok:      true,
		Message: "Reservation success",
	}, w)
}

type CreateReservationResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

type ValidationErrorResponse struct {
	Ok             bool    `json:"ok"`
	Message        string  `json:"message"`
	NotValidFields []Field `json:"not_valid_fields"`
}
type Field struct {
	Name  string
	Value string
}

func (r *ReservationController) CheckAvailableHandler(w http.ResponseWriter, req *http.Request) {
	r.Templater.Create(w, req, "check-available.page.tmpl", &templateCreator.TemplateData{})
}
