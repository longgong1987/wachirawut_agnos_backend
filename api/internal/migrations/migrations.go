package migrations

import (
	"fmt"
	"wachirawut_agnos_backend/internal/repositories"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&repositories.StrongPasswordStep{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Migrated successfully!")
}
