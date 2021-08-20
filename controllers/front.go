package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"goapi.com/api/controllers/book"
	"goapi.com/api/controllers/point"
)

var (
	router = gin.Default()
)

func RegisterControllers() {
	router.GET("/books", book.GetAll)

	router.GET("/points", point.GetAll)
	router.GET("/points/:id", point.Get)
	router.POST("/points", point.Post)
	router.PUT("/points/:id", point.Put)
	router.DELETE("/points:id", point.Delete)
}

func Start(params ...int) {
	port := 3000
	if len(params) > 0 {
		port = params[0]
	}

	router.Run(fmt.Sprintf("%s:%d", "localhost", port))
}
