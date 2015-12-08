/*
--- Part Two ---

Now, given the same instructions, find the position of the first character that causes him to enter the basement (floor -1). The first character in the instructions has position 1, the second character has position 2, and so on.

For example:

) causes him to enter the basement at character position 1.
()()) causes him to enter the basement at character position 5.
What is the position of the character that causes Santa to first enter the basement?*/
package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
)

func main() {
	log.Println("AdventOfCode.com - Day 1 - Puzzle 2")

	var err error
	var character byte

	floor := 0
	position := 0

	fileBytes, err := ioutil.ReadFile("./puzzle-input.txt")
	if err != nil {
		log.Fatalf("Error reading puzzle file: %s", err.Error())
	}

	reader := bytes.NewReader(fileBytes)

	for {
		character, err = reader.ReadByte()

		if err == nil {
			position++

			if string(character) == "(" {
				floor++
			}

			if string(character) == ")" {
				floor--
			}
		}

		if err == io.EOF {
			break
		}

		if floor == -1 {
			break
		}
	}

	log.Printf("Santa reached floor -1 at position %d", position)
}
