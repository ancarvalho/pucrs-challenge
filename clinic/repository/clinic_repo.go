package repository

import (
	"clinic/types"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ClinicRepository struct {
	Db *sqlx.DB
}

func InstanciateClinicRepo(db *sqlx.DB) *ClinicRepository {
	return &ClinicRepository{Db: db}
}

func (r *ClinicRepository) GetAppointments() ([]types.Appointment, error) {

	appointments := []types.Appointment{}

	err := r.Db.Select(&appointments, `
		SELECT 
			a.time, 
			a.id, 
			a.user_id, 
			u.name as user_name 
		FROM appointments a 
		LEFT JOIN users u 
		ON u.id = a.user_id
	
	`)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error retrieving appointments")
	}
	return appointments, nil

}

func (r *ClinicRepository) GetCustomers() ([]types.User, error) {

	users := []types.User{}

	err := r.Db.Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil, errors.New("error retrieving users")
	}
	return users, nil

}

func (r *ClinicRepository) CreateUser(user *types.User) error {

	createUser := ` INSERT INTO users (id, name, phone) VALUES (?, ?, ?)`

	res, err := r.Db.MustExec(createUser, user.Id, user.Name, user.Phone).RowsAffected()

	if err != nil {
		return errors.New("error creating user")
	}

	if res > 0 {
		fmt.Println("User Created")
	}

	return nil

}

func (r *ClinicRepository) MakeAppointment(appointment types.Appointment) error {

	createUser := ` INSERT INTO appointments ( id, user_id, time) VALUES (?, ?, ?)`

	res, err := r.Db.MustExec(createUser, appointment.Id, appointment.UserID, appointment.Time).RowsAffected()

	if err != nil {
		return errors.New("error creating appointment")
	}

	if res > 0 {
		fmt.Println("Appointment Created")
	}

	return nil

}

func (r *ClinicRepository) CancelAppointment(appointment_id uuid.UUID) error {

	var err error

	cancelAppointment := ` DELETE FROM appointments WHERE id = ?`

	res, err := r.Db.MustExec(cancelAppointment, appointment_id).RowsAffected()

	if err != nil {
		return errors.New("error canceling appointment")
	}

	if res > 0 {
		fmt.Println("Appointment Canceled")
	}

	return nil
}
