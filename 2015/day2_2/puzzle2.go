/*
--- Part Two ---

The elves are also running low on ribbon. Ribbon is all the same width, so they only have to worry about the length they need to order, which they would again like to be exact.

The ribbon required to wrap a present is the shortest distance around its sides, or the smallest perimeter of any one face. Each present also requires a bow made out of ribbon as well; the feet of ribbon required for the perfect bow is equal to the cubic feet of volume of the present. Don't ask how they tie the bow, though; they'll never tell.

For example:

A present with dimensions 2x3x4 requires 2+2+3+3 = 10 feet of ribbon to wrap the present plus 2*3*4 = 24 feet of ribbon for the bow, for a total of 34 feet.
A present with dimensions 1x1x10 requires 1+1+1+1 = 4 feet of ribbon to wrap the present plus 1*1*10 = 10 feet of ribbon for the bow, for a total of 14 feet.
How many total feet of ribbon should they order?
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
	log.Println("AdventOfCode.com - Day 2 - Puzzle 2")

	var err error

	file, err := os.Open("./puzzle-input.txt")
	if err != nil {
		log.Fatalf("Error reading puzzle file: %s", err.Error())
	}

	defer file.Close()

	reader := bufio.NewScanner(file)
	line := ""
	lineIndex := 0

	totalRibbonFeet := 0

	var box *Box

	for reader.Scan() {
		lineIndex++

		line = reader.Text()
		log.Printf("Analyzing line '%s' (%d)...\n", line, lineIndex)

		box = parseLine(line)
		totalRibbonFeet += calculateRibbonLength(box)
	}

	log.Printf("The elves should order %d feet of ribbon\n", totalRibbonFeet)
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

func calculateRibbonLength(box *Box) int {
	return (box.Length * 2) + (box.Width * 2) + (box.Length * box.Width * box.Height)
}
