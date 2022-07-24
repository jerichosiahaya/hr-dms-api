package services

import (
	"hr-dms-api/models"
	"hr-dms-api/repositories"
)

type Service interface {
	CreateEmployee(employee *models.Employee) (*models.Employee, error)
}

type service struct {
	repository repositories.Repository
}

func NewController(repository *repositories.Repository) service {
	return service{repository: *repository}
}

func (s *service) CreateEmployee(input *models.Employee) (*models.Employee, error) {
	var employee *models.Employee
	employee.First_name = input.First_name
	employee.Last_name = input.Last_name
	employee.Phone_number = input.Phone_number
	employee.Hire_date = input.Hire_date
	employee.Job_id = input.Job_id
	employee.Office_id = input.Office_id
	employee.Photo = input.Photo
	// res, err := repositories.Repository.Save(employee)
	// if err != nil {
	// 	return res, err
	// }
	return employee, nil
}
