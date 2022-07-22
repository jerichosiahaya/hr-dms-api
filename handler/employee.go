// this is the gate

package handler

import (
	"hr-dms-api/employee"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	employeeService employee.Service
}

func NewEmployeeHandler(employeeService employee.Service) *EmployeeHandler {
	return &EmployeeHandler{employeeService}
}

func (h *EmployeeHandler) GetAllEmployees(c *gin.Context) {
	employees, err := h.employeeService.GetEmployees()
	if err != nil {
		errorMessage := gin.H{"errors": err.Error}
		c.JSON(http.StatusNotFound, errorMessage)
	}
	c.JSON(http.StatusOK, employees)
}

func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	test := gin.H{"m": "j"}
	c.JSON(http.StatusOK, test)
}
