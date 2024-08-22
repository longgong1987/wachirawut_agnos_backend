package repositories

import "gorm.io/gorm"

type Repository struct {
	postgresql *gorm.DB
}

func NewRepository(postgresql *gorm.DB) *Repository {
	return &Repository{
		postgresql: postgresql,
	}
}
