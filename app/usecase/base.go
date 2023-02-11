package usecase

import "github.com/emerfauzan/cake-store-api/app/repository"

type Usecase struct {
	repository *repository.Repository
}

func Init(repository repository.Repository) *Usecase {
	return &Usecase{
		repository: &repository,
	}
}
