/*
--- Part Two ---

Now, take the signal you got on wire a, override wire b to that signal, and reset the other wires (including wire a). What new signal is ultimately provided to wire a?
*/

package main

import (
	"bufio"
	"log"
	"os"

	"github.com/adampresley/lexer"
)

func main() {
	log.Println("AdventOfCode.com - Day 7 - Puzzle 2")

	var wires WireCollection
	var wire *Wire
	var instruction string

	instructions, err := os.Open("./puzzle-input.txt")
	if err != nil {
		log.Fatalf("Error reading instructions file: %s", err.Error())
		os.Exit(1)
	}

	defer instructions.Close()

	reader := bufio.NewScanner(instructions)
	reader.Split(bufio.ScanLines)
	wires = NewWireCollection()

	for reader.Scan() {
		instruction = reader.Text()

		wire = NewWire(instruction)
		wires[wire.Name] = wire
	}

	/*
	 * Calculate wire 'a'
	 */
	wireAResult := wires.EvaluateWireValue("a")

	log.Printf("Wire 'a' first result == %d\n", wireAResult)

	/*
	 * Now, clear our wire cache, replace the source for wire 'b'
	 * with the result of 'a', and re-calc.
	 */
	wires.ClearCache()
	wires["b"].Source = make([]lexer.Token, 0)

	wires["b"].Source = append(wires["b"].Source, lexer.Token{
		Type:  TOKEN_NUMBER,
		Value: wireAResult,
	})

	log.Printf("Wire 'a' second result == %d\n", wires.EvaluateWireValue("a"))

}
