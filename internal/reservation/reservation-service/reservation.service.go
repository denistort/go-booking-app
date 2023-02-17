package reservation_service

import (
	"github.com/denistort/go-booking-app/internal/reservation/reservation-entity"
)

type ReservationService struct {
}

func New() *ReservationService {
	return &ReservationService{}
}

func (r *ReservationService) CreateReservation(
	firstName string,
	lastName string,
	phone string,
	email string,
	dateFrom string,
	dateTo string) *reservation_entity.ReservationEntity {
	return reservation_entity.New(firstName, lastName, phone, email, dateFrom, dateTo)
}
