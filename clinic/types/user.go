package types

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type User struct {
	Id    uuid.UUID `db:"id" validate:"required"`
	Name  string    `db:"name" validate:"required,min=3,max=40"`
	Phone int       `db:"phone" validate:"min=8"`
}

func ValidateUser(name string, phone int) (*User, error) {
	user := User{Id: uuid.New(), Name: name, Phone: phone}
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error validating user")
	}
	return &user, nil

}
