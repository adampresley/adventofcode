package main

import "log"

type GPS struct {
	Location Point
}

func NewGPS() *GPS {
	return &GPS{
		Location: NewPoint(0, 0),
	}
}

func (gps *GPS) MoveUp() {
	log.Printf("Moving up...")
	gps.Location.Y++
}

func (gps *GPS) MoveDown() {
	log.Printf("Moving down...")
	gps.Location.Y--
}

func (gps *GPS) MoveLeft() {
	log.Printf("Moving left...")
	gps.Location.X--
}

func (gps *GPS) MoveRight() {
	log.Printf("Moving right...")
	gps.Location.X++
}
