package employee

type Employee struct {
	Id           int
	First_name   string
	Last_name    string
	Phone_number string
	Hire_date    string
	Office_id    int
	Job_id       int
	Photo        string
}

type Employees []Employee

type InputEmployee struct {
	First_name   string `json:"first_name" binding:"required"`
	Last_name    string `json:"last_name" binding:"required"`
	Phone_number string `json:"phone_number" binding:"required"`
	Hire_date    string `json:"hire_date" binding:"required"`
	Office_id    int    `json:"office_id" binding:"required"`
	Job_id       int    `json:"job_id" binding:"required"`
	Photo        string `json:"photo" binding:"required"`
}
