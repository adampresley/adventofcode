package main

import (
	"fmt"
	"log"
	"strings"
)

type StringParser struct {
	input  string
	length int
}

func NewStringParser() *StringParser {
	return &StringParser{
		input:  "",
		length: 0,
	}
}

func (parser *StringParser) DetermineStatus(input string) StringStatus {
	parser.input = input
	parser.length = len(input)

	msg := fmt.Sprintf("Determining if %s is naughty or nice...", input)

	if parser.isNice() {
		log.Printf("%s NICE!", msg)
		return NICE
	}

	log.Printf("%s NAUGHTY!", msg)
	return NAUGHTY
}

func (parser *StringParser) isNice() bool {
	repeatingPairsCount := 0
	sandwichLettersCount := 0

	pair := EOL

	for index, char := range parser.input {
		pair = parser.getPair(index)
		if pair != EOL {
			if parser.isRepeatingPair(pair) {
				repeatingPairsCount++
			}
		}

		if parser.isSandwich(index, string(char)) {
			sandwichLettersCount++
		}

		if repeatingPairsCount >= MINIMUM_REQUIRED_REPEATERS && sandwichLettersCount >= MINIMUM_REQUIRED_SANDWICHES {
			return true
		}
	}

	return false
}

func (parser *StringParser) getPair(currentIndex int) string {
	if currentIndex+1 < parser.length {
		return parser.input[currentIndex : currentIndex+2]
	}

	return EOL
}

func (parser *StringParser) isRepeatingPair(pair string) bool {
	return strings.Count(parser.input, pair) > 1
}

func (parser *StringParser) isSandwich(currentIndex int, currentChar string) bool {
	if currentIndex+2 < parser.length {
		sample := parser.input[currentIndex : currentIndex+3]
		if sample[2:3] == currentChar {
			return true
		}
	}

	return false
}
