package reservation_entity

type ReservationEntity struct {
	FirstName string
	LastName  string
	Phone     string
	Email     string
	DateFrom  string
	DateTo    string
}

func New(firstName string, lastName string, phone string, email string, dateFrom string, dateTo string) *ReservationEntity {

	return &ReservationEntity{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		Email:     email,
		DateTo:    dateTo,
		DateFrom:  dateFrom,
	}
}
