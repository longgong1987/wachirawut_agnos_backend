package repositories

import (
	"wachirawut_agnos_backend/internal/domain"

	"gorm.io/gorm"
)

type Repositories struct {
	Postgresql *gorm.DB
}

func NewRepository(postgresql *gorm.DB) *Repositories {
	return &Repositories{
		Postgresql: postgresql,
	}
}

func (repo *Repositories) CreateStrongPasswordStep(strongPassword *domain.StrongPasswordStepDtO) error {

	// create strong password step
	strongPasswordStep := &StrongPasswordStep{
		Password: strongPassword.Password,
		Step:     strongPassword.Step,
	}

	if err := repo.Postgresql.Create(strongPasswordStep).Error; err != nil {
		return err
	}

	return nil
}
