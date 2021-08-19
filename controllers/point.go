package controllers

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"goapi.com/api/common"
	"goapi.com/api/models"
)

type PointController struct {
	PointIDPattern *regexp.Regexp
}

const (
	route = "points"
)

var (
	simpleRoute   = "/" + route
	composedRoute = "/" + route + "/"
)

func (pc PointController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(strings.ToLower(r.URL.Path), " ")
	if strings.EqualFold(path, simpleRoute) {
		switch r.Method {
		case http.MethodGet:
			pc.GetAll(w)
		case http.MethodPost:
			pc.Post(w, r)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	if strings.Index(path, composedRoute) == 0 {
		matches := pc.PointIDPattern.FindStringSubmatch(path)
		if len(matches) != 2 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		switch r.Method {
		case http.MethodGet:
			pc.Get(id, w)
		case http.MethodPut:
			pc.Put(id, w, r)
		case http.MethodDelete:
			pc.Delete(id, w, r)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

// constructor is a convention for a func
func NewPointController() *PointController {
	return &PointController{
		PointIDPattern: regexp.MustCompile(`^/points/(\d+)/?$`),
	}
}

func (pc *PointController) GetAll(w http.ResponseWriter) {
	_, _, fn := common.Trace()
	fmt.Println(fn)

	err := common.WriteJson(models.GetPoints(), w)
	if err != nil {
		common.WriteError(err, w)
	}
}

func (pc *PointController) Get(id int, w http.ResponseWriter) {
	_, _, fn := common.Trace()
	fmt.Println(fn)

	point, err1 := models.GetPointById(id)
	if err1 != nil {
		common.WriteError(err1, w)
		return
	}

	err2 := common.WriteJson(point, w)
	if err2 != nil {
		common.WriteError(err2, w)
	}
}

func (pc *PointController) Post(w http.ResponseWriter, r *http.Request) {
	_, _, fn := common.Trace()
	fmt.Println(fn)

	var newPoint models.Point
	err1 := common.ReadJson(&newPoint, w, r)
	if err1 != nil {
		common.WriteError(err1, w)
		return
	}

	createdPoint, err2 := models.AddPoint(newPoint)
	if err2 != nil {
		common.WriteError(err2, w)
		return
	}

	err3 := common.WriteJsonWithStatusCode(createdPoint, http.StatusCreated, w)
	if err3 != nil {
		common.WriteError(err3, w)
	}
}

func (pc *PointController) Put(id int, w http.ResponseWriter, r *http.Request) {
	_, _, fn := common.Trace()
	fmt.Println(fn)

	var existingPoint models.Point
	err1 := common.ReadJson(&existingPoint, w, r)
	if err1 != nil {
		common.WriteError(err1, w)
		return
	}

	if id != existingPoint.ID {
		common.WriteErrorWithStatusCode(fmt.Errorf("ID of submitted point must match ID in URL"), http.StatusBadRequest, w)
		return
	}

	createdPoint, err2 := models.UpdatePoint(existingPoint)
	if err2 != nil {
		common.WriteError(err2, w)
		return
	}

	err3 := common.WriteJson(createdPoint, w)
	if err3 != nil {
		common.WriteError(err3, w)
	}
}

func (pc *PointController) Delete(id int, w http.ResponseWriter, r *http.Request) {
	_, _, fn := common.Trace()
	fmt.Println(fn)

	var existingPoint models.Point
	err1 := common.ReadJson(&existingPoint, w, r)
	if err1 != nil {
		common.WriteError(err1, w)
		return
	}

	if id != existingPoint.ID {
		common.WriteErrorWithStatusCode(fmt.Errorf("ID of submitted point must match ID in URL"), http.StatusBadRequest, w)
		return
	}

	err2 := models.DeletePoint(existingPoint)
	if err2 != nil {
		common.WriteError(err2, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
