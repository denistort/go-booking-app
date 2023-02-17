package JsonController

import (
	"encoding/json"
	"fmt"
	"github.com/denistort/go-booking-app/cmd/web/config"
	"net/http"
)

// JsonController is handlers type
type JsonController struct {
	AppConfig *config.AppConfig
}

// New creates the JsonController
func New(config *config.AppConfig) *JsonController {
	return &JsonController{
		AppConfig: config,
	}
}

func (j *JsonController) ContactJsonHandler(w http.ResponseWriter, req *http.Request) {
	dateFrom := req.Form.Get("date-from")
	dateTo := req.Form.Get("date-to")

	fmt.Println(req.Form.Encode())
	fmt.Println(dateFrom, dateTo)
	resp := PostReservationResponse{
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

type PostReservationResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}
