package main

import "fmt"

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) Point {
	return Point{
		X: x,
		Y: y,
	}
}

func (point Point) String() string {
	return fmt.Sprintf("%d:%d", point.X, point.Y)
}
