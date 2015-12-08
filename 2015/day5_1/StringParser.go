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

var naughtyBits = []string{
	"ab",
	"cd",
	"pq",
	"xy",
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

	if parser.hasNaughtyBits() {
		log.Printf("%s NAUGHTY!", msg)
		return NAUGHTY
	}

	if parser.isNice() {
		log.Printf("%s NICE!", msg)
		return NICE
	}

	return NAUGHTY
}

func (parser *StringParser) hasNaughtyBits() bool {
	for _, naughtyBit := range naughtyBits {
		if strings.Contains(parser.input, naughtyBit) {
			return true
		}
	}

	return false
}

func (parser *StringParser) isNice() bool {
	vowelCount := 0
	sequentialLetterCount := 0

	for index, char := range parser.input {
		if parser.isVowel(string(char)) {
			vowelCount++
		}

		if parser.isDoubleDigit(string(char), index) {
			sequentialLetterCount++
		}

		if vowelCount >= MINIMUM_REQUIRED_VOWELS && sequentialLetterCount >= MINIMUM_REQUIRED_DOUBLE_DIGITS {
			return true
		}
	}

	return false
}

func (parser *StringParser) isVowel(char string) bool {
	return strings.ContainsAny(char, "aeiou")
}

func (parser *StringParser) isDoubleDigit(currentChar string, currentPosition int) bool {
	if currentPosition+1 < parser.length {
		expected := fmt.Sprintf("%s%s", currentChar, currentChar)
		actual := parser.input[currentPosition : currentPosition+2]

		return expected == actual
	}

	return false
}
