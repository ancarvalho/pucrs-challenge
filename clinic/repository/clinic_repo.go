package repository

import (
	"clinic/types"
	"errors"
)

func Get_appointments() {

}

func Create_user(name string, phone int) error {
	e, user := types.ValidateUser(name, phone)
	if e != nil {
		return errors.New("Error Validating User")
	}

}

func Make_appointment() {

}

func Cancel_appointment() {

}
