package models

import (
	"goapi.com/api/data"
	"gorm.io/gorm"
)

type DBPoint struct {
	gorm.Model
	X float32
	Y float32
	Z float32
}

var (
	ctx = data.GetContext()
)

func Migrate() {
	ctx.AutoMigrate(&DBPoint{})
}

func Create(point DBPoint) {
	ctx.Create(&point)
}

func Save(point DBPoint) {
	ctx.Save(&point)
}

func First(id int) DBPoint {
	var point DBPoint
	ctx.First(&point, id)
	return point
}
