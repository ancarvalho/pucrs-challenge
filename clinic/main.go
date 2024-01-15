package main

import (
	"clinic/db"
	"clinic/repository"
	"clinic/usecases"
	"fmt"
	"os"
)

func main() {

	db, err := db.InstanciateDB()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	repo := repository.InstanciateClinicRepo(db)
	usecases.InstanciateUsecase(repo).Home()

	// usecases.Home()
}
