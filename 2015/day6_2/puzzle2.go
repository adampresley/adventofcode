/*
--- Part Two ---

You just finish implementing your winning light pattern when you realize you mistranslated Santa's message from Ancient Nordic Elvish.

The light grid you bought actually has individual brightness controls; each light can have a brightness of zero or more. The lights all start at zero.

The phrase turn on actually means that you should increase the brightness of those lights by 1.

The phrase turn off actually means that you should decrease the brightness of those lights by 1, to a minimum of zero.

The phrase toggle actually means that you should increase the brightness of those lights by 2.

What is the total brightness of all lights combined after following Santa's instructions?

For example:

turn on 0,0 through 0,0 would increase the total brightness by 1.
toggle 0,0 through 999,999 would increase the total brightness by 2000000.
*/

package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/adampresley/lexer"
)

func main() {
	log.Println("AdventOfCode.com - Day 6 - Puzzle 1")

	grid := &[1000][1000]int{}
	var token lexer.Token

	var action *Action
	var coordinateIndex int
	var coordinates [4]int

	instructions, err := ioutil.ReadFile("./puzzle-input.txt")
	if err != nil {
		log.Fatalf("Error reading instructions file: %s", err.Error())
		os.Exit(1)
	}

	lex := lexer.NewLexer("instructions", string(instructions), LexerBegin)
	lex.Run()

	for {
		_ = "breakpoint"
		token = lex.NextToken()

		if token.Type == lexer.TOKEN_ERROR {
			log.Println(token.Value)
			break
		}

		/*
		 * If we have a command, create a new Action structure. Zero out the coordinate index.
		 * We are going to start tracking the instruction.
		 */
		if token.Type == TOKEN_COMMAND {
			if coordinateIndex == 4 {
				action.StartX = coordinates[0]
				action.StartY = coordinates[1]
				action.EndX = coordinates[2]
				action.EndY = coordinates[3]

				action.Execute(grid)
			}

			action = &Action{}
			coordinateIndex = 0

			switch strings.TrimSpace((token.Value).(string)) {
			case COMMAND_TOGGLE:
				action.Fn = Toggle

			case COMMAND_TURN_OFF:
				action.Fn = TurnOff

			case COMMAND_TURN_ON:
				action.Fn = TurnOn
			}
		}

		if token.Type == TOKEN_NUMBER {
			coordinates[coordinateIndex], _ = strconv.Atoi((token.Value).(string))
			coordinateIndex++
		}

		if token.Type == lexer.TOKEN_EOF {
			action.StartX = coordinates[0]
			action.StartY = coordinates[1]
			action.EndX = coordinates[2]
			action.EndY = coordinates[3]

			action.Execute(grid)
			break
		}
	}

	brightness := 0

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			brightness += grid[x][y]
		}
	}

	log.Printf("Brightness is %d", brightness)
}
