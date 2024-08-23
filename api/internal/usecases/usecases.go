package usecases

import (
	"wachirawut_agnos_backend/internal/domain"
	"wachirawut_agnos_backend/internal/repositories"
)

type Usecase interface {
	CheckStrongPasswordStep(passwordStep *domain.StrongPasswordStep) error
}

type usecases struct {
	Repo repositories.Repositories
}

func NewUsecase(repo repositories.Repositories) Usecase {
	return &usecases{
		Repo: repo,
	}
}

func (uc *usecases) CheckStrongPasswordStep(passwordStep *domain.StrongPasswordStep) error {
	return nil
}
