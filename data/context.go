package data

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	context *gorm.DB
)

func NewContext() {
	dsn := "root:GwVE7uc2eYT73FNG@tcp(127.0.0.1:3307)/goapi?charset=utf8mb4&parseTime=True&loc=Local"
	context, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	if context == nil {
		panic(fmt.Errorf("Context could not be created"))
	}
}

func GetContext() *gorm.DB {
	return context
}
