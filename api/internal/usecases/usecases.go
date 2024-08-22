package usecases

import "wachirawut_agnos_backend/internal/repositories"

type Usecase struct {
}

func New(repo *repositories.Repository) *Usecase {
	return &Usecase{}
}
