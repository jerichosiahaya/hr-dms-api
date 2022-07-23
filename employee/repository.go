package employee

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(employee Employee) (Employee, error)
	Get() (Employees, error)
	GetById(Id int) (Employee, error)
	// Delete(Id int)
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

func (r *repository) GetById(Id int) (Employee, error) {
	var employee Employee
	err := r.db.Where("id = ?", Id).Find(&employee).Error
	if err != nil {
		return employee, err
	}
	return employee, nil
}
