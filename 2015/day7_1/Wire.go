package main

import (
	"strconv"

	"github.com/adampresley/lexer"
)

type Wire struct {
	Name   string
	Source []lexer.Token
}

func NewWire(instruction string) *Wire {
	lex := lexer.NewLexer("instruction", instruction, LexerBegin)
	lex.Run()

	result := &Wire{
		Source: make([]lexer.Token, 0),
	}

	var token lexer.Token

	for {
		token = lex.NextToken()

		if token.Type == lexer.TOKEN_EOF {
			break
		}

		if token.Type == TOKEN_CONNECTOR {
			token = lex.NextToken()
			result.Name = (token.Value).(string)
			break
		}

		result.Source = append(result.Source, token)
	}

	return result
}

func (wire *Wire) Evaluate(wires WireCollection) int {
	_ = "breakpoint" // Wire.Evaluate

	var result1 int
	var result2 int
	var result3 int

	/*
	 * Number or wire
	 */
	if len(wire.Source) == 1 {
		// Number
		if wire.Source[0].Type == TOKEN_NUMBER {
			result1, _ = strconv.Atoi((wire.Source[0].Value).(string))
			return result1
		}

		// Wire
		return wires.EvaluateWireValue((wire.Source[0].Value).(string))
	}

	/*
	 * NOT gate
	 */
	if len(wire.Source) == 2 {
		if wire.Source[1].Type == TOKEN_NUMBER {
			result1, _ = strconv.Atoi((wire.Source[1].Value).(string))
			return result1
		}

		result1 = wires.EvaluateWireValue((wire.Source[1].Value).(string))
		return int(^uint16(result1))
	}

	/*
	 * Other GATE expressions
	 */
	if wire.Source[0].Type == TOKEN_NUMBER {
		result1, _ = strconv.Atoi((wire.Source[0].Value).(string))
	} else {
		result1 = wires.EvaluateWireValue((wire.Source[0].Value).(string))
	}

	if wire.Source[2].Type == TOKEN_NUMBER {
		result2, _ = strconv.Atoi((wire.Source[2].Value).(string))
	} else {
		result2 = wires.EvaluateWireValue((wire.Source[2].Value).(string))
	}

	switch (wire.Source[1].Value).(string) {
	case "AND":
		result3 = result1 & result2

	case "OR":
		result3 = result1 | result2

	case "LSHIFT":
		result3 = int(uint32(result1) << uint32(result2))

	case "RSHIFT":
		result3 = int(uint32(result1) >> uint32(result2))
	}

	return result3
}
