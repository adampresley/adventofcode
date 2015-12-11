package main

import (
	"github.com/adampresley/lexer"
)

const (
	TOKEN_WIRE lexer.TokenType = iota
	TOKEN_NUMBER
	TOKEN_GATE
	TOKEN_CONNECTOR

	CONNECTOR   string = "->"
	GATE_AND    string = "AND"
	GATE_OR     string = "OR"
	GATE_LSHIFT string = "LSHIFT"
	GATE_RSHIFT string = "RSHIFT"
	GATE_NOT    string = "NOT"
)
