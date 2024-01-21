package models

import (
	"errors"
	"gorm.io/gorm"
)

func MigrateModel(db *gorm.DB) error {
	// AutoMigrate creates or updates tables based on the model structs
	if err := db.AutoMigrate(&Customer{}, &UserRole{}, &Product{}, &Order{}, &HistoryCart{}); err != nil {
		return err
	}

	// Check if the UserRole with ID 1 exists
	var userRole UserRole
	if result := db.First(&userRole, 1); result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// UserRole with ID 1 doesn't exist, insert the initial data
		if err := db.Create(&UserRole{ID: 1, Role: "customer"}).Error; err != nil {
			return err
		}
	}

	return nil
}
