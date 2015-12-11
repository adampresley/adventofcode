package main

import (
	"strings"

	"github.com/adampresley/lexer"
)

func LexerBegin(lex *lexer.Lexer) lexer.LexFn {
	_ = "breakpoint" // LexerBegin
	lex.SkipWhitespace()

	if lex.IsEOF() {
		return nil
	}

	// WIRE, GATE, or NUMBER
	if lex.IsNumber() {
		return LexerNumber
	}

	if beginNextIsGate(lex) {
		return LexerGate
	}

	return LexerWire
}

func beginNextIsGate(lex *lexer.Lexer) bool {
	var gates = []string{
		GATE_AND,
		GATE_OR,
		GATE_LSHIFT,
		GATE_RSHIFT,
		GATE_NOT,
	}

	for _, gate := range gates {
		if strings.HasPrefix(lex.InputToEnd(), gate) {
			return true
		}
	}

	return false
}
