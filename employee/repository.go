package employee

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(employee Employee) (Employee, error)
	Get() (Employees, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(employee Employee) (Employee, error) {
	err := r.db.Create(&employee).Error
	if err != nil {
		return employee, err
	}
	return employee, nil
}

func (r *repository) Get() (Employees, error) {
	var employees Employees
	err := r.db.Find(&employees).Error
	if err != nil {
		return employees, err
	}
	return employees, nil
}
