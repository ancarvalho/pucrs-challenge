package utils

func Interpret_command(command int, curr_page *int, total_pages int, items_per_page int) {

	if command == 2 && *curr_page > 0 {
		*curr_page -= 1

	}
	if command == 9 && *curr_page < total_pages {
		*curr_page += 1

	}

	// idx := Transform_index_back(command, *curr_page, items_per_page)
	// fmt.Println(idx)

	// names := utils.GetNames()

	// fmt.Println("Curr Index", idx, names[idx])

}

func Transform_index_back(command int, curr_page int, items_per_page int) int {
	return command + (curr_page * items_per_page) - 3
}

func Convert_to_index(curr_page int, items_per_row int, index int) int {
	return index - (curr_page * items_per_row) + 3
}

func Determine_pages(items_per_row int, total_len int) (int, int) {
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

func Transform(curr_page int, items_per_row int, total_pages int, last_page_len int) (int, int) {
	start_index := curr_page * items_per_row

	last_index := start_index

	if curr_page == (total_pages - 1) {
		last_index += last_page_len

	} else {
		last_index += items_per_row
	}

	return start_index, last_index
}
