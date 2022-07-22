package main

import (
	"fmt"
	"hr-dms-api/controllers"
	"hr-dms-api/repositories"
	"log"

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

	Repository := repositories.NewRepository(db)
	Controller := controllers.NewController(Repository)

}
