package employee

import (
	"errors"
	"time"
)

type Service interface {
	CreateEmployee(input InputEmployee) (Employee, error)
	GetEmployees() (Employees, error)
	GetEmployeeById(input GetEmployeeDetailById) (Employee, error)
	DeleteEmployee(input GetEmployeeDetailById) (Employee, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateEmployee(input InputEmployee) (Employee, error) {
	employee := Employee{}
	employee.First_name = input.First_name
	employee.Last_name = input.Last_name
	employee.Phone_number = input.Phone_number
	employee.Hire_date = input.Hire_date
	employee.Job_id = input.Job_id
	employee.Office_id = input.Office_id
	employee.Photo = input.Photo
	res, err := s.repository.Save(employee)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *service) GetEmployees() (Employees, error) {
	tempEmployees := Employees{}
	tempEmployees, err := s.repository.Get()
	if err != nil {
		return tempEmployees, err
	}
	employees := Employees{}
	tempEmployee := Employee{}
	for _, s := range tempEmployees {
		tempEmployee.First_name = s.First_name
		tempEmployee.Last_name = s.Last_name
		tempEmployee.Id = s.Id
		tempEmployee.Photo = s.Photo
		tempEmployee.Office_id = s.Office_id
		tempEmployee.Job_id = s.Job_id
		tempEmployee.Phone_number = s.Phone_number

		newTime, _ := time.Parse("2006-01-02T00:00:00+09:00", s.Hire_date)
		tempEmployee.Hire_date = newTime.Format("2006-01-02")
		employees = append(employees, tempEmployee)
	}
	return employees, nil
}

// get employee by id
func (s *service) GetEmployeeById(input GetEmployeeDetailById) (Employee, error) {
	res, err := s.repository.GetById(input.Id)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *service) DeleteEmployee(input GetEmployeeDetailById) (Employee, error) {
	res, err := s.repository.GetById(input.Id)
	if err != nil {
		return res, err
	}

	if res.Id != input.Id {
		return res, errors.New("owner validation failed")
	}

	deleteResponse, err := s.repository.Delete(res.Id)
	if err != nil {
		return deleteResponse, err
	}
	return deleteResponse, nil
}
