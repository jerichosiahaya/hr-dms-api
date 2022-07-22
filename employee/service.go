package employee

type Service interface {
	CreateEmployee(input InputEmployee) (Employee, error)
	GetEmployees() (Employees, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateEmployee(input InputEmployee) (Employee, error) {
	var employee Employee
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
	employee := Employees{}
	employee, err := s.repository.Get()
	if err != nil {
		return employee, err
	}
	return employee, nil
}
