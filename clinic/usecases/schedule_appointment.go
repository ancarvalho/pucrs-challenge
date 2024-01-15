package usecases

import (
	"clinic/utils"
	"fmt"
	"os"
	"strconv"
)

func (u *Usecase) ScheduleAppointment() {

	items_per_page := 6

	var input string
	var err error
	var command int

	customers, err := u.ClinicRepo.GetCustomers()
	if err != nil || len(customers) < 1 {
		fmt.Println("Error Getting Users")
		u.Home()
		os.Exit(0)
	}

	total_lenght := len(customers)
	total_pages, last_page_len := utils.Determine_pages(items_per_page, total_lenght)
	var curr_page int = 0

	for {

		fmt.Println("Press 1 to Return Home")
		if curr_page > 0 {
			fmt.Println("Press 2 Previous Page")
		}

		start_index, last_index := utils.Transform(curr_page, items_per_page, total_pages, last_page_len)
		// fmt.Println(start_index, last_index)
		for i := start_index; i < last_index; i++ {
			idx := utils.Convert_to_index(curr_page, items_per_page, i)
			fmt.Println("Press", idx, "to Schedule an Appointment for", customers[i].Name)
		}

		// fmt.Println(total_pages)

		if curr_page < total_pages {
			fmt.Println("Press 9 to Next Page")
		}

		fmt.Println("Press 0 to Exit")

		fmt.Scanln(&input)
		command, err = strconv.Atoi(input)
		if len(input) == 1 && err == nil && (command == 0 || command == 1) {
			break
		}

		if len(input) == 1 && err == nil && (command == 2 || command == 9) {
			utils.Interpret_command(command, &curr_page, total_pages, items_per_page)
			continue
		}

		if len(input) == 1 && err == nil && (command >= 3 || command < 9) {
			break
		}

		fmt.Println("Value Invalid")
		continue

	}

	if command <= 1 && err == nil {
		u.DefaultRoutes(command)
	}

	idx := utils.Transform_index_back(command, curr_page, items_per_page)
	u.MakeAppointment(customers[idx].Id)
	os.Exit(0)

}
