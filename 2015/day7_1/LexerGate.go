package main

import "github.com/adampresley/lexer"

/*
LexerGate processes logic gate operators.
*/
func LexerGate(lex *lexer.Lexer) lexer.LexFn {
	_ = "breakpoint" // LexerGate
	lex.SkipWhitespace()

	for {
		if lex.IsEOF() {
			lex.Errorf("Unexpected EOF after a gate. Expected a wire")
			return LexerEnd
		}

		if lex.IsWhitespace() {
			lex.Emit(TOKEN_GATE)
			return LexerBegin
		}

		lex.Inc(1)
	}
}
