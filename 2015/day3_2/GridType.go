package main

import (
	"fmt"
	"log"
)

type GridType map[Point]int

/*
Visits a location. Returns the number of times this house has been visited
*/
func (grid GridType) Visit(x, y int) int {
	point := NewPoint(x, y)
	message := fmt.Sprintf("Dropping presents off at %s. This house now has", point)

	if count, ok := grid[point]; ok {
		count++
		grid[point] = count

		message = fmt.Sprintf("%s %d presents", message, count)
		log.Println(message)

		return count
	}

	message = fmt.Sprintf("%s 1 present", message)
	log.Println(message)

	grid[point] = 1
	return 1
}
