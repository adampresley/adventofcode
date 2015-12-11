package main

import (
	"strings"

	"github.com/adampresley/lexer"
)

/*
LexerNumber processes numbers. Numbers in this dialect represent coordinates.
*/
func LexerNumber(lex *lexer.Lexer) lexer.LexFn {
	_ = "breakpoint" // LexerNumber
	lex.SkipWhitespace()

	for {
		if lex.IsEOF() {
			lex.Errorf("Unexpected EOF after a number. Expected either a connector or gate")
			return LexerEnd
		}

		if lex.IsNumber() {
			lex.Inc()
			continue
		}

		/*
		 * Next up can be either GATE or CONNECTOR
		 */
		nextFn := whatIsAfterNumber(lex)

		lex.Emit(TOKEN_NUMBER)
		return nextFn
	}
}

func whatIsAfterNumber(lex *lexer.Lexer) lexer.LexFn {
	nextCharacters := lex.PeekCharacters(3)
	if strings.Contains(nextCharacters, CONNECTOR) {
		return LexerConnector
	}

	return LexerGate
}
