package main

import (
	"strings"

	"github.com/adampresley/lexer"
)

/*
LexerConnector processes the connector operator ->. The connector operator
is used to assign, or connect a value to a WIRE
*/
func LexerConnector(lex *lexer.Lexer) lexer.LexFn {
	_ = "breakpoint" // LexerConnector
	lex.SkipWhitespace()

	if !strings.Contains(lex.InputToEnd(), CONNECTOR) {
		return lex.Errorf("A connector was expected but not found")
	}

	lex.Inc(2)

	lex.Emit(TOKEN_CONNECTOR)
	return LexerWire
}
