package types

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Appointment struct {
	Id       uuid.UUID `db:"id" validate:"required"`
	Time     time.Time `db:"time" validate:"required"`
	UserID   uuid.UUID `db:"user_id" validate:"required"`
	UserName string    `db:"user_name"`
}

func ValidateAppointment(user_id uuid.UUID, schedule_time time.Time) (*Appointment, error) {
	var err error

	now := time.Now()

	if schedule_time.Before(now) {
		return nil, errors.New("error validating data, it can`t be at past")
	}

	appointment := Appointment{Id: uuid.New(), Time: schedule_time, UserID: user_id}
	validate := validator.New()
	err = validate.Struct(appointment)
	if err != nil {

		return nil, errors.New("error validating appointment")
	}
	return &appointment, nil

}
