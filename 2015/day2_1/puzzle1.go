/*
--- Day 2: I Was Told There Would Be No Math ---

The elves are running low on wrapping paper, and so they need to submit an order for more. They have a list of the dimensions (length l, width w, and height h) of each present, and only want to order exactly as much as they need.

Fortunately, every present is a box (a perfect right rectangular prism), which makes calculating the required wrapping paper for each gift a little easier: find the surface area of the box, which is 2*l*w + 2*w*h + 2*h*l. The elves also need a little extra paper for each present: the area of the smallest side.

For example:

A present with dimensions 2x3x4 requires 2*6 + 2*12 + 2*8 = 52 square feet of wrapping paper plus 6 square feet of slack, for a total of 58 square feet.
A present with dimensions 1x1x10 requires 2*1 + 2*10 + 2*10 = 42 square feet of wrapping paper plus 1 square foot of slack, for a total of 43 square feet.
All numbers in the elves' list are in feet. How many total square feet of wrapping paper should they order?
*/
package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Box struct {
	Length int
	Width  int
	Height int
}

func main() {
	log.Println("AdventOfCode.com - Day 2 - Puzzle 1")

	var err error

	file, err := os.Open("./puzzle-input.txt")
	if err != nil {
		log.Fatalf("Error reading puzzle file: %s", err.Error())
	}

	defer file.Close()

	reader := bufio.NewScanner(file)
	line := ""
	totalSquareFeet := 0
	lineIndex := 0

	var box *Box
	var area int
	var slack int

	for reader.Scan() {
		lineIndex++

		line = reader.Text()
		log.Printf("Analyzing line '%s' (%d)...\n", line, lineIndex)

		box = parseLine(line)
		area = calculateBoxSquareFeet(box)
		slack = calculateSlackFeet(box)

		totalSquareFeet = totalSquareFeet + (area + slack)
	}

	log.Printf("The elves should order %d square feet of wrapping paper\n", totalSquareFeet)
}

func parseLine(line string) *Box {
	stringValues := strings.Split(line, "x")
	result := make([]int, 3)

	for index, value := range stringValues {
		result[index], _ = strconv.Atoi(value)
	}

	sort.Ints(result)

	return &Box{
		Length: result[0],
		Width:  result[1],
		Height: result[2],
	}
}

func calculateBoxSquareFeet(box *Box) int {
	return 2 * ((box.Length * box.Width) + (box.Width * box.Height) + (box.Height * box.Length))
}

func calculateSlackFeet(box *Box) int {
	sizes := []int{box.Length, box.Width, box.Height}
	sort.Ints(sizes)

	return sizes[0] * sizes[1]
}
