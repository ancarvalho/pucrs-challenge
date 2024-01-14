package usecases

import (
	"clinic/utils"
	"fmt"
	"strconv"
)

func ScheduleAppointment() {

	items_per_page := 6

	var input string
	var err error
	var command int

	names := utils.GetNames()
	total_lenght := len(names)
	total_pages, last_page_len := determine_pages(items_per_page, total_lenght)
	var curr_page int = 0

	for {

		fmt.Println("Press 1 to Return Home")
		if curr_page > 0 {
			fmt.Println("Press 2 Previous Page")
		}

		start_index, last_index := transform(curr_page, items_per_page, total_pages, last_page_len)
		fmt.Println(start_index, last_index)
		for i := start_index; i < last_index; i++ {
			idx := convert_to_index(curr_page, items_per_page, i)
			fmt.Println("Press", idx, "to Schedule an Appointment for", names[i])
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

		if len(input) == 1 && err == nil && (command >= 2 || command <= 9) {
			interpret_command(command, &curr_page, total_pages, items_per_page)
			continue
		}

		fmt.Println("Value Invalid")
		continue

	}

	IntermediaryCommand(command)

}

func interpret_command(command int, curr_page *int, total_pages int, items_per_page int) {

	if command == 2 && *curr_page > 0 {
		*curr_page -= 1

	}
	if command == 9 && *curr_page < total_pages {
		*curr_page += 1

	}

	idx := transform_index_back(command, *curr_page, items_per_page)
	fmt.Println(idx)

	// names := utils.GetNames()

	// fmt.Println("Curr Index", idx, names[idx])

}

func transform_index_back(command int, curr_page int, items_per_page int) int {
	return command + (curr_page * items_per_page) - 3
}

func convert_to_index(curr_page int, items_per_row int, index int) int {
	return index - (curr_page * items_per_row) + 3
}

func determine_pages(items_per_row int, total_len int) (int, int) {
	var last_page_len int

	total_pages := total_len / items_per_row

	if remainig := total_len % items_per_row; remainig != 0 {
		total_pages += 1
		last_page_len = remainig
	}
	if last_page_len <= 0 {
		last_page_len = items_per_row
	}

	return total_pages, last_page_len

}

func transform(curr_page int, items_per_row int, total_pages int, last_page_len int) (int, int) {
	start_index := curr_page * items_per_row

	last_index := start_index

	if curr_page == (total_pages - 1) {
		last_index += last_page_len

	} else {
		last_index += items_per_row
	}

	return start_index, last_index
}
