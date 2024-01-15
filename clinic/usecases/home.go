package usecases

import (
	"fmt"
	"os"
	"strconv"
)

func (u *Usecase) DefaultRoutes(command int) {
	switch command {
	case 0:
		// panic("Exiting...")
		fmt.Println("Exiting...")
		os.Exit(0)
	case 1:
		u.Home()
		// panic("Exiting")
		fmt.Println("Exiting...")
		os.Exit(0)
	}
}

func (u *Usecase) Home() {
	var input string
	var err error
	var val int

	fmt.Println("Press 1 to Create User")
	fmt.Println("Press 2 to Make Appointment")
	fmt.Println("Press 3 to Cancel Appointment")
	fmt.Println("Press 0 to Exit")

	for {
		fmt.Scanln(&input)
		val, err = strconv.Atoi(input)
		if err == nil && val >= 0 && val <= 3 {
			fmt.Println("Valid integer input")
			break
		} else {
			fmt.Println("Invalid integer input")
			continue
		}
	}

	switch val {
	case 0:
		fmt.Println("Saindo...")

	case 1:
		u.CreateUser()
	case 2:
		u.ScheduleAppointment()
	case 3:
		u.CancelAppointment()

	}
}
