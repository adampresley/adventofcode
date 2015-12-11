package main

import (
	"strings"

	"github.com/adampresley/lexer"
)

/*
LexerWire processes a wire. A wire is like a variable. It holds values
*/
func LexerWire(lex *lexer.Lexer) lexer.LexFn {
	_ = "breakpoint" // LexerWire
	lex.SkipWhitespace()

	for {
		if lex.IsEOF() {
			lex.Emit(TOKEN_WIRE)
			return LexerEnd
		}

		if lex.IsNewline() {
			lex.Emit(TOKEN_WIRE)
			return LexerBegin
		}

		if lex.IsWhitespace() {
			lex.Emit(TOKEN_WIRE)

			nextCharacters := lex.PeekCharacters(3)
			if strings.Contains(nextCharacters, CONNECTOR) {
				return LexerConnector
			}

			return LexerGate
		}

		lex.Inc(1)
	}
}
