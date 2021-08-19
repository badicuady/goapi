package models

import (
	"fmt"
)

type Point struct {
	ID int
	X  float32
	Y  float32
	Z  float32
}

var (
	points []*Point = []*Point{}
	ID              = 1
)

func GetPoints() []*Point {
	return points
}

func AddPoint(p Point) (Point, error) {
	p.ID = ID
	ID++
	points = append(points, &p)
	return p, nil
}

func GetPointById(id int) (Point, error) {
	for _, v := range points {
		if v.ID == id {
			return *v, nil
		}
	}
	return Point{}, fmt.Errorf("Point with ID '%v' not found", id)
}

func UpdatePoint(p Point) (Point, error) {
	for i, v := range points {
		if v.ID == p.ID {
			points[i] = &p
			return p, nil
		}
	}
	return Point{}, fmt.Errorf("Point with ID '%v' not found", p.ID)
}

func DeletePoint(p Point) error {
	for i, v := range points {
		if v.ID == p.ID {
			points = append(points[:i], points[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Point with ID '%v' not found", p.ID)
}
