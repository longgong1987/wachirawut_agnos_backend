package usecases

import (
	"wachirawut_agnos_backend/internal/domain"
	"wachirawut_agnos_backend/internal/repositories"
)

type Usecase interface {
	CheckStrongPasswordStep(passwordStep *domain.StrongPasswordStepDtO) (domain.StrongPasswordStepResponse, error)
}

type usecase struct {
	Repo repositories.Repositories
}

func NewUsecase(repo repositories.Repositories) *usecase {
	return &usecase{
		Repo: repo,
	}
}

func (uc *usecase) CheckStrongPasswordStep(passwordStep *domain.StrongPasswordStepDtO) (domain.StrongPasswordStepResponse, error) {

	var numOfSteps domain.StrongPasswordStepResponse

	if err := uc.Repo.CreateStrongPasswordStep(passwordStep); err != nil {
		return numOfSteps, err
	}

	numOfSteps.NumOfSteps = 1

	return numOfSteps, nil
}
