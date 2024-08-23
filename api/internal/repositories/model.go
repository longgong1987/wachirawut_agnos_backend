package repositories

import (
	"time"

	"gorm.io/gorm"
)

type StrongPasswordStep struct {
	ID        int        `gorm:"column:id; primarykey; autoIncrement:true" json:"id"`
	Password  string     `json:"password"`
	Step      int        `json:"step"`
	CreatedBy string     `gorm:"column:created_by" json:"created_by"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedBy string     `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedBy string     `gorm:"column:deleted_by" json:"deleted_by"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (e *StrongPasswordStep) TableName() string {
	return "strong_password_steps"
}

func (u *StrongPasswordStep) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	return nil
}
