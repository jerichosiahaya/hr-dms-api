package employee

type Employee struct {
	Id           int    `json:"id"`
	First_name   string `json:"first_name"`
	Last_name    string `json:"last_name"`
	Phone_number int    `json:"phone_number"`
	Hire_date    string `json:"hire_date"`
	Office_id    int    `json:"office_id"`
	Job_id       int    `json:"job_id"`
	Photo        string `json:"photo"`
}

type Employees []Employee

type InputEmployee struct {
	First_name   string `json:"first_name" binding:"required"`
	Last_name    string `json:"last_name" binding:"required"`
	Phone_number int    `json:"phone_number" binding:"required"`
	Hire_date    string `json:"hire_date"`
	Office_id    int    `json:"office_id"`
	Job_id       int    `json:"job_id"`
	Photo        string `json:"photo"`
}

type GetEmployeeDetailById struct {
	Id int `uri:"id" binding:"required"`
}
