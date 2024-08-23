package repositories

import (
	"wachirawut_agnos_backend/internal/domain"

	"gorm.io/gorm"
)

type Repositories interface {
	CreateStrongPasswordStep(strongPassword *domain.StrongPasswordStep) error
}

type Repository struct {
	Postgresql gorm.DB
}

func NewRepository(postgresql gorm.DB) Repositories {
	return &Repository{
		Postgresql: postgresql,
	}
}

func (repo *Repository) CreateStrongPasswordStep(strongPassword *domain.StrongPasswordStep) error {
	return nil
}
