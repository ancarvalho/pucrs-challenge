package usecases

import (
	"clinic/types"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func (u *Usecase) CreateUser() {
	var input, name string
	var err, command_err error
	var matched bool
	// var name string
	var phone int
	command := 999

	reg, err := regexp.Compile(`^[a-zA-Z]+$`)

	if err != nil {
		fmt.Println("Regex Error", err)
	}

	fmt.Println("Type Customer Name:")
	for {
		fmt.Scanln(&input)

		fmt.Println("Length", len(input))

		if command, command_err = strconv.Atoi(input); len(input) == 1 && err == nil && (command == 0 || command == 1) {
			break
		}
		matched = reg.MatchString(input)

		if matched {
			fmt.Println("Valid input")
			name = input
			break
		} else {
			fmt.Println("Invalid input")
			continue
		}
	}
	// fmt.Println(command)
	if command <= 1 && command_err == nil {
		u.DefaultRoutes(command)
	}

	fmt.Println("Type Customer Phone:")
	for {
		fmt.Scanln(&input)

		if command, command_err = strconv.Atoi(input); len(input) == 1 && err == nil && (command == 0 || command == 1) {
			break
		}

		phone, err = strconv.Atoi(input)

		if err == nil && len(input) >= 8 {
			fmt.Println("Valid Phone Number")
			break
		}

		continue
	}

	if command <= 1 && command_err == nil {
		// fmt.Println("here", command)
		u.DefaultRoutes(command)
	}

	user, err := types.ValidateUser(name, phone)
	if err != nil {
		fmt.Println(err)
		u.CreateUser()
		os.Exit(0)
	}

	err = u.ClinicRepo.CreateUser(user)
	if err != nil {
		fmt.Println(err)
		u.CreateUser()
		os.Exit(0)
	}

	u.Home()
	os.Exit(0)

}
