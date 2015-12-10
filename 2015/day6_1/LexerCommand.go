package main

import "github.com/adampresley/lexer"

/*
LexerCommand processes a command text. Valid commands are:

toggle
turn on
turn off
*/
func LexerCommand(lex *lexer.Lexer) lexer.LexFn {
	_ = "breakpoint"
	lex.SkipWhitespace()

	for {
		if lex.IsEOF() {
			return lex.Errorf("Unexpected EOF")
		}

		if lex.IsNumber() {
			lex.Emit(TOKEN_COMMAND)
			return LexerNumber
		}

		lex.Inc()
	}
}
