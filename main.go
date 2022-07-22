package main

import (
	"fmt"
	"hr-dms-api/employee"
	"hr-dms-api/handler"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(localhost)/dimas_hrm_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Conection to database is good")

	router := gin.Default()
	api := router.Group("/api/v1")

	// employee
	employeeRepository := employee.NewRepository(db)
	employeeService := employee.NewService(employeeRepository)
	employeeHandler := handler.NewEmployeeHandler(employeeService)
	api.GET("/employees", employeeHandler.GetAllEmployees)
	api.POST("/employees", employeeHandler.CreateEmployee)

	router.Run()

}
