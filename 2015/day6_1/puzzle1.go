/*
--- Day 6: Probably a Fire Hazard ---

Because your neighbors keep defeating you in the holiday house decorating contest year after year, you've decided to deploy one million lights in a 1000x1000 grid.

Furthermore, because you've been especially nice this year, Santa has mailed you instructions on how to display the ideal lighting configuration.

Lights in your grid are numbered from 0 to 999 in each direction; the lights at each corner are at 0,0, 0,999, 999,999, and 999,0. The instructions include whether to turn on, turn off, or toggle various inclusive ranges given as coordinate pairs. Each coordinate pair represents opposite corners of a rectangle, inclusive; a coordinate pair like 0,0 through 2,2 therefore refers to 9 lights in a 3x3 square. The lights all start turned off.

To defeat your neighbors this year, all you have to do is set up your lights by doing the instructions Santa sent you in order.

For example:

* turn on 0,0 through 999,999 would turn on (or leave on) every light.
* toggle 0,0 through 999,0 would toggle the first line of 1000 lights, turning off the ones that were on, and turning on the ones that were off.
* turn off 499,499 through 500,500 would turn off (or leave off) the middle four lights.

After following the instructions, how many lights are lit?
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

	numLightsLit := 0

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] == 1 {
				numLightsLit++
			}
		}
	}

	log.Printf("%d lights are lit", numLightsLit)
}
