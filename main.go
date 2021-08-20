package main

import (
	"fmt"

	"goapi.com/api/controllers"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dsn := "root:GwVE7uc2eYT73FNG@tcp(127.0.0.1:3307)/goapi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&Product{})

	Start()
}

func Start(params ...int) {
	fmt.Println("Starting...")

	port := 3000
	if len(params) > 0 {
		port = params[0]
	}

	controllers.RegisterControllers()
	controllers.Start(port)
}
