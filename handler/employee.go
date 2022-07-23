// this is the gate

package handler

import (
	"hr-dms-api/employee"
	"hr-dms-api/helper"
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
		// errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("failed to fetch the data", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIResponse("successfully fetch the data", http.StatusOK, "success", employees)
	c.JSON(http.StatusOK, response)
}

func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	var input employee.InputEmployee
	err := c.ShouldBindJSON(&input)
	if err != nil {
		// errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("register employee failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	res, err := h.employeeService.CreateEmployee(input)
	if err != nil {
		response := helper.APIResponse("register employee failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("employee registered successfully", http.StatusOK, "success", res)
	c.JSON(http.StatusOK, response)
}

// delete employee by id
func (h *EmployeeHandler) DeleteEmployee(c *gin.Context) {
	var input employee.GetEmployeeDetailById

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("failed to delete employee", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	res, err := h.employeeService.DeleteEmployee(input)
	if err != nil {
		response := helper.APIResponse("failed to delete employee", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("successfully deleted the employee", http.StatusOK, "success", res)
	c.JSON(http.StatusOK, response)

}
