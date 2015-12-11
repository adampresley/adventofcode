package main

import "github.com/adampresley/lexer"

/*
LexerNumber processes numbers. Numbers in this dialect represent coordinates.
*/
func LexerNumber(lex *lexer.Lexer) lexer.LexFn {
	_ = "breakpoint"
	lex.SkipWhitespace()

	for {
		if lex.IsEOF() {
			lex.Emit(TOKEN_NUMBER)
			return LexerEnd
		}

		if lex.IsNumber() {
			lex.Inc(1)
			continue
		}

		lex.Emit(TOKEN_NUMBER)
		return LexerDelimiter
	}
}
