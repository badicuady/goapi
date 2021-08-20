package point

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"goapi.com/api/common"
	"goapi.com/api/models/point"
)

func GetAll(c *gin.Context) {
	_, _, fn := common.Trace()
	fmt.Println(fn)

	points := point.GetPoints()
	c.IndentedJSON(http.StatusOK, points)
}

func Get(c *gin.Context) {
	_, _, fn := common.Trace()
	fmt.Println(fn)

	id, err1 := strconv.Atoi(c.Param("id"))
	if err1 != nil {
		c.IndentedJSON(http.StatusInternalServerError, err1)
		return
	}

	point, err2 := point.GetPointById(id)
	if err2 != nil {
		c.IndentedJSON(http.StatusInternalServerError, err2)
		return
	}

	c.IndentedJSON(http.StatusOK, point)
}

func Post(c *gin.Context) {
	_, _, fn := common.Trace()
	fmt.Println(fn)

	var newPoint point.Point
	err1 := common.ReadJson(&newPoint, c.Request.Body)
	if err1 != nil {
		c.IndentedJSON(http.StatusInternalServerError, err1)
		return
	}

	createdPoint, err2 := point.AddPoint(newPoint)
	if err2 != nil {
		c.IndentedJSON(http.StatusInternalServerError, err2)
		return
	}

	c.IndentedJSON(http.StatusCreated, createdPoint)
}

func Put(c *gin.Context) {
	_, _, fn := common.Trace()
	fmt.Println(fn)

	var existingPoint point.Point
	err1 := common.ReadJson(&existingPoint, c.Request.Body)
	if err1 != nil {
		c.IndentedJSON(http.StatusInternalServerError, err1)
		return
	}

	id, err2 := strconv.Atoi(c.Param("id"))
	if err2 != nil {
		c.IndentedJSON(http.StatusInternalServerError, err2)
		return
	}

	if id != existingPoint.ID {
		c.IndentedJSON(http.StatusBadRequest, fmt.Errorf("ID of submitted point must match ID in URL"))
		return
	}

	createdPoint, err3 := point.UpdatePoint(existingPoint)
	if err3 != nil {
		c.IndentedJSON(http.StatusInternalServerError, err3)
		return
	}

	c.IndentedJSON(http.StatusOK, createdPoint)
}

func Delete(c *gin.Context) {
	_, _, fn := common.Trace()
	fmt.Println(fn)

	var existingPoint point.Point
	err1 := common.ReadJson(&existingPoint, c.Request.Body)
	if err1 != nil {
		c.IndentedJSON(http.StatusInternalServerError, err1)
		return
	}

	id, err2 := strconv.Atoi(c.Param("id"))
	if err2 != nil {
		c.IndentedJSON(http.StatusInternalServerError, err2)
		return
	}

	if id != existingPoint.ID {
		c.IndentedJSON(http.StatusBadRequest, fmt.Errorf("ID of submitted point must match ID in URL"))
		return
	}

	err3 := point.DeletePoint(existingPoint)
	if err3 != nil {
		c.IndentedJSON(http.StatusInternalServerError, err3)
		return
	}

	c.Status(http.StatusContinue)
}
