package reservation_dto

type ReservationDto struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Email     string `validate:"required,email"`
	Phone     string `validate:"required"`
	DateFrom  string `validate:"required"`
	DateTo    string `validate:"required"`
}
