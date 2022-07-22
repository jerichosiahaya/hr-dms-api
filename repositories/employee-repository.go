package repositories

import (
	"hr-dms-api/models"

	"gorm.io/gorm"
)

type Repository interface {
	Save(employee *models.Employee) (*models.Employee, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(employee *models.Employee) (*models.Employee, error) {
	err := r.db.Create(&employee).Error
	if err != nil {
		return employee, err
	}
	return employee, nil
}
