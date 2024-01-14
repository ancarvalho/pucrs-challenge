package usecases

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func IntermediaryCommand(command int) {

	switch command {
	case 0:
		// panic("Exiting...")
		fmt.Println("Exiting...")
		os.Exit(0)
	case 1:
		Home()
		// panic("Exiting")
		fmt.Println("Exiting...")
		os.Exit(0)
	}
}

func CreateUser() {
	var input, name string
	var err error
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

		if command, err = strconv.Atoi(input); len(input) == 1 && err == nil && (command == 0 || command == 1) {
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
	if command <= 1 {
		IntermediaryCommand(command)
	}

	fmt.Println("Type Customer Phone:")
	for {
		fmt.Scanln(&input)

		if command, err := strconv.Atoi(input); len(input) == 1 && err == nil && (command == 0 || command == 1) {
			break
		}

		phone, err = strconv.Atoi(input)

		if err == nil && len(input) >= 8 {
			fmt.Println("Valid Phone Number")
			break
		}

		continue
	}

	if command <= 1 {
		IntermediaryCommand(command)
	}

	fmt.Println("Name:", name, "Phone:", phone)

}
