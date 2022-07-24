package main

import (
	"fmt"
	"hr-dms-api/employee"
	"hr-dms-api/handler"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	SetupServer()
}

func SetupServer() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api/v1")
	// server.Use(middlewares.CORSMiddleware())
	// server.Use(sentrygin.New(sentrygin.Options{}))
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"serverTime": time.Now().UTC().Unix(),
			"status":     "system is working fine",
		})
	})

	dsn := "root:@tcp(localhost)/dimas_hrm_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("conection to database is good")

	// employee endpoint
	employeeRepository := employee.NewRepository(db)
	employeeService := employee.NewService(employeeRepository)
	employeeHandler := handler.NewEmployeeHandler(employeeService)
	api.GET("/employees", employeeHandler.GetAllEmployees)
	api.GET("/employees/:id", employeeHandler.GetOneEmployee)
	api.POST("/employees", employeeHandler.CreateEmployee)
	api.DELETE("/employees/:id", employeeHandler.DeleteEmployee)
	api.PUT("/employees/:id", employeeHandler.UpdateEmployee)

	router.Run()

	return router
}
