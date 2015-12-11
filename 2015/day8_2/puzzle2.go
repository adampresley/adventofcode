/*
--- Part Two ---

Now, let's go the other way. In addition to finding the number of characters of code, you should now encode each code representation as a new string and find the number of characters of the new encoded representation, including the surrounding double quotes.

For example:

"" encodes to "\"\"", an increase from 2 characters to 6.
"abc" encodes to "\"abc\"", an increase from 5 characters to 9.
"aaa\"aaa" encodes to "\"aaa\\\"aaa\"", an increase from 10 characters to 16.
"\x27" encodes to "\"\\x27\"", an increase from 6 characters to 11.
Your task is to find the total number of characters to represent the newly encoded strings minus the number of characters of code in each original string literal. For example, for the strings above, the total encoded length (6 + 9 + 16 + 11 = 42) minus the characters in the original code representation (23, just like in the first part of this puzzle) is 42 - 23 = 19.
*/

package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/adampresley/lexer"
)

func main() {
	log.Println("AdventOfCode.com - Day 8 - Puzzle 2")

	santaList, err := ioutil.ReadFile("./puzzle-input.txt")
	if err != nil {
		log.Fatalf("Error reading Santa's list file: %s", err.Error())
		os.Exit(1)
	}

	lex := lexer.NewLexer("Santa's List", string(santaList), LexerBegin)
	lex.Run()

	var token lexer.Token
	var stringLiteralCount int
	var stringMemoryCount int
	var difference int

	for {
		token = lex.NextToken()

		if token.IsEOF() {
			break
		}

		if token.IsError() {
			log.Println(token.String())
			break
		}

		stringLiteralCount += len(token.String())
		stringMemoryCount += memoryFootprint(token)
	}

	difference = stringMemoryCount - stringLiteralCount
	log.Printf("Literal character count: %d, Memory count: %d\n", stringLiteralCount, stringMemoryCount)
	log.Printf("Difference: %d\n", difference)
}

func memoryFootprint(token lexer.Token) int {
	if token.Type == TOKEN_NEWLINE {
		return 0
	}

	/*
	 * 1 for the quote. One for an opposite open/close quote. 1 for escape character
	 */
	if token.Type == TOKEN_QUOTE {
		return 3
	}

	if token.Type == TOKEN_ESCAPED_CHARACTER {
		return len(token.String()) + 2
	}

	if token.Type == TOKEN_HEX_CODE {
		return len(token.String()) + 1
	}

	return len(token.String())
}
