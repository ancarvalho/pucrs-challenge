package types

import (
	"github.com/go-playground/validator/v10"
)

type User struct {
	Name  string `json:"name" validate:"required,min=3,max=40"`
	Phone int    `json:"phone" validate:"min=8, max=16"`
}

func ValidateUser(name string, phone int) (validator.ValidationErrors, *User) {
	user := User{Name: name, Phone: phone}
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		return errors, nil
	}
	return nil, &user

}
