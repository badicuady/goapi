package main

import (
	"fmt"
	"net/http"

	"goapi.com/api/controllers"
)

func main() {
	Start()
}

func Start(params ...int) {
	fmt.Println("Starting...")

	port := 3000
	if len(params) > 0 {
		port = params[0]
	}

	controllers.RegisterControllers()
	fmt.Printf("Server started on port: %v!\r\n", port)

	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
