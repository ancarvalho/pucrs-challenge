package usecases

import "clinic/repository"

type Usecase struct {
	ClinicRepo *repository.ClinicRepository
}

func InstanciateUsecase(clinicRepo *repository.ClinicRepository) *Usecase {
	return &Usecase{ClinicRepo: clinicRepo}
}
