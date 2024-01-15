package usecases

import (
	"clinic/types"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func (u *Usecase) MakeAppointment(user_id uuid.UUID) {

	var input, input_data, input_time string
	var err, command_err error
	var matched bool
	var command = 999

	var now = time.Now()
	var y, _, d = now.Date()
	var m = int(now.Month())

	date_reg, err := regexp.Compile(`^(0[1-9]|[12][0-9]|3[01])\/(0[1-9]|1[012])\/(20)\d{2}$`)

	if err != nil {
		fmt.Println("Regex Error", err)
	}

	fmt.Println("Type Date (DD/MM/YYYY):")
	for {
		fmt.Scanln(&input)

		// fmt.Println("Length", len(input))

		if command, command_err = strconv.Atoi(input); len(input) == 1 && err == nil && (command == 0 || command == 1) {
			break
		}
		matched = date_reg.MatchString(input)

		if matched {
			var date []int
			dataInput := strings.Split(input, "/")
			for i := 0; i < len(dataInput); i++ {
				v, _ := strconv.Atoi(dataInput[i])
				date = append(date, v)
			}

			fmt.Println(m)
			if date[0] >= d && date[1] >= m && date[2] >= y {
				input_data = input
				fmt.Println("Valid input")
				break
			}
		} else {
			fmt.Println("Invalid input")
			continue
		}
	}
	if command <= 1 && command_err == nil {
		fmt.Println(err)
		u.DefaultRoutes(command)
	}

	time_reg, err := regexp.Compile(`^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$`)
	if err != nil {
		fmt.Println("Regex Error", err)
	}

	fmt.Println("Type Time (20:00):")

	for {
		fmt.Scanln(&input)

		// fmt.Println("Length", len(input))

		if command, command_err = strconv.Atoi(input); len(input) == 1 && err == nil && (command == 0 || command == 1) {
			break
		}
		matched = time_reg.MatchString(input)

		if matched {
			fmt.Println("Valid input")
			input_time = input
			break
		} else {
			fmt.Println("Invalid input")
			continue
		}
	}

	if command <= 1 && command_err == nil {
		u.DefaultRoutes(command)
	}

	// fmt.Println(transformToTime(input_data, input_time))

	time, err := transformToTime(input_data, input_time)
	if err != nil {
		fmt.Println(err)
		u.MakeAppointment(user_id)
		os.Exit(0)
	}

	var appointment *types.Appointment

	appointment, err = types.ValidateAppointment(user_id, time)
	if err != nil {
		fmt.Println(err)
		u.Home()
		os.Exit(0)
	}

	u.ClinicRepo.MakeAppointment(*appointment)

	u.Home()
	os.Exit(0)

}

func transformToTime(d string, t string) (time.Time, error) {
	const DDMMYYYYHHMM = "02/01/2006 15:04"

	return time.Parse(DDMMYYYYHHMM, d+" "+t)

}
