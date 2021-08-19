package controllers

import "net/http"

func RegisterControllers() {
	pc := NewPointController()

	http.Handle("/points", *pc)
	http.Handle("/points/", *pc)
}
